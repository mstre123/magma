// Copyright 2021 The Magma Authors.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"context"
	"os"

	"github.com/magma/magma/src/go/protos/magma/capture"
	"github.com/magma/magma/src/go/protos/magma/mconfig"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

//go:generate go run github.com/golang/mock/mockgen -source=middleware.go -package mock_middleware -destination mock_middleware/mock_middleware_test.go Middleware

// middleware defines middlewares functions - used for generating mocks.
type middleware interface {
	GetUnaryClientInterceptor() grpc.DialOption
	withPassthroughUnaryClientInterceptor() grpc.DialOption
	withCaptureUnaryClientInterceptor() grpc.DialOption
}

type Middleware struct {
	writePath string
}

// NewMiddleware returns a configured middleware.
func NewMiddleware(c *mconfig.MiddlewareConfig) Middleware {
	return Middleware{
		writePath: c.GetInterceptOutputPath(),
	}
}

// GetUnaryClientInterceptor returns a UnaryClientInterceptor based on the middlewares configed options.
func (m *Middleware) GetUnaryClientInterceptor() grpc.DialOption {
	if m.writePath == "" {
		return m.withPassthroughUnaryClientInterceptor()
	}
	return m.withCaptureUnaryClientInterceptor()
}

// withPassthroughUnaryClientInterceptor returns an interceptor that is a passthrough and takes no action on the intercepted rpc calls.
func (m *Middleware) withPassthroughUnaryClientInterceptor() grpc.DialOption {
	passthroughInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	return grpc.WithUnaryInterceptor(passthroughInterceptor)
}

// withCaptureUnaryClientInterceptor returns an interceptor that captures and writes all rpc calls that go through it a file.
func (m *Middleware) withCaptureUnaryClientInterceptor() grpc.DialOption {

	captureInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		rq, err := anypb.New(req.(proto.Message))
		if err != nil {
			panic(err)
		}
		rp, err := anypb.New(reply.(proto.Message))
		if err != nil {
			panic(err)
		}

		cap := capture.UnaryCall{
			Method:   method,
			Request:  rq,
			Response: rp,
			Err:      "",
		}

		if err := write(cap, m.writePath); err != nil {
			panic(err)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}

	return grpc.WithUnaryInterceptor(captureInterceptor)
}

// write is a helper function that writes intercepted rpc calls to a file.
func write(call capture.UnaryCall, path string) error {
	// /home/vagrant/magma/src/go/capture/resources/intercept.txt
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(call.String())); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
