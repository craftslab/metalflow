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

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/craftslab/metalflow/model"
	"github.com/craftslab/metalflow/util"
)

// GetServerVersion godoc
// @Summary Get server version
// @Description Get server version
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {string} model.Version
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /config/server/version [get]
func (c *Controller) GetServerVersion(ctx *gin.Context) {
	version, err := model.ServerVersion()
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, version)
}
