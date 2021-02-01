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
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/craftslab/metalflow/model"
	"github.com/craftslab/metalflow/util"
)

// GetAccount godoc
// @Summary Get account by ID
// @Description Get account by ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path uint true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /accounts/{id} [get]
func (c *controller) GetAccount(ctx *gin.Context) {
	param := ctx.Param("id")

	if param == "self" {
		if account, err := model.GetSelfAccount(); err == nil {
			ctx.JSON(http.StatusOK, account)
		} else {
			util.NewError(ctx, http.StatusNotFound, err)
		}
	} else {
		if id, err := strconv.ParseUint(param, 10, 64); err == nil {
			if account, e := model.GetAccount(uint(id)); e == nil {
				ctx.JSON(http.StatusOK, account)
			} else {
				util.NewError(ctx, http.StatusNotFound, e)
			}
		} else {
			util.NewError(ctx, http.StatusBadRequest, err)
		}
	}
}

// QueryAccount godoc
// @Summary Query account
// @Description Query account
// @Tags accounts
// @Accept json
// @Produce json
// @Param q query string true "Username search by q"
// @Success 200 {object} model.Account
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /accounts [get]
func (c *controller) QueryAccount(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")

	account, err := model.QueryAccount(q)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}
