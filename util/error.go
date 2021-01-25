// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(ctx *gin.Context, status int, err error) {
	e := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	ctx.JSON(status, e)
}
