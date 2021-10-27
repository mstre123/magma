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
	"reflect"
	"testing"

	"github.com/magma/magma/src/go/protos/magma/mconfig"
	"github.com/stretchr/testify/assert"
)

func TestNewMiddleware_Success(t *testing.T) {
	path := "path"
	cfg := &mconfig.MiddlewareConfig{
		InterceptOutputPath: path,
	}
	mw := NewMiddleware(cfg)
	assert.Equal(t, path, mw.writePath)
}

func TestNewMiddleware_EmptyConfig(t *testing.T) {
	path := ""
	mw := NewMiddleware(nil)
	assert.Equal(t, path, mw.writePath)
}

func TestMiddleware_GetUnaryClientInterceptor_Passthrough(t *testing.T) {
	path := ""
	cfg := &mconfig.MiddlewareConfig{
		InterceptOutputPath: path,
	}
	mw := NewMiddleware(cfg)
	interceptor := mw.GetUnaryClientInterceptor()
	assert.Equal(t, "*grpc.funcDialOption", reflect.TypeOf(interceptor).String())
}

func TestMiddleware_GetUnaryClientInterceptor_Capture(t *testing.T) {
	path := "path"
	cfg := &mconfig.MiddlewareConfig{
		InterceptOutputPath: path,
	}
	mw := NewMiddleware(cfg)
	interceptor := mw.GetUnaryClientInterceptor()
	assert.Equal(t, "*grpc.funcDialOption", reflect.TypeOf(interceptor).String())
}
