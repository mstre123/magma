/*
 Copyright 2020 The Magma Authors.

 This source code is licensed under the BSD-style license found in the
 LICENSE file in the root directory of this source tree.

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package swagger

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ValidateLteAuthOp validates the  lte_auth_op is 16 bytes.
func ValidateLteAuthOp(formats strfmt.Registry, lteAuthOp strfmt.Base64) error {

	if err := validate.Required("lte_auth_op", "body", strfmt.Base64(lteAuthOp)); err != nil {
		return err
	}

	if err := validate.MinLength("lte_auth_op", "body", string(lteAuthOp), 15); err != nil {
		return err
	}

	if err := validate.MaxLength("lte_auth_op", "body", string(lteAuthOp), 16); err != nil {
		return err
	}

	return nil
}
