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

// GetNode godoc
// @Summary Get node by ID
// @Description Get node by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path uint true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id} [get]
func (c *controller) GetNode(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
	}

	node, err := model.GetNode(uint(id))
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, node)
}

// GetHealth godoc
// @Summary Get node health by ID
// @Description Get node health by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path uint true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id}/health [get]
func (c *controller) GetHealth(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
	}

	health, err := model.GetHealth(uint(id))
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, health)
}

// GetInfo godoc
// @Summary Get node information by ID
// @Description Get node information by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path uint true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id}/info [get]
func (c *controller) GetInfo(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
	}

	info, err := model.GetInfo(uint(id))
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, info)
}

// GetPerf godoc
// @Summary Get node performance by ID
// @Description Get node performance by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path uint true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id}/perf [get]
func (c *controller) GetPerf(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
	}

	perf, err := model.GetPerf(uint(id))
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, perf)
}

// QueryNode godoc
// @Summary Query node
// @Description Query node
// @Tags nodes
// @Accept json
// @Produce json
// @Param q query string true "ID search by q"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes [get]
func (c *controller) QueryNode(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")

	node, err := model.QueryNode(q)
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, node)
}

// AddNode godoc
// @Summary Add node
// @Description Add node
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path uint true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes [put]
func (c *controller) AddNode(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
	}

	node, err := model.AddNode(uint(id))
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, node)
}

// DelNode godoc
// @Summary Delete node
// @Description Delete node
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path uint true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes [delete]
func (c *controller) DelNode(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		util.NewError(ctx, http.StatusBadRequest, err)
	}

	node, err := model.DelNode(uint(id))
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, node)
}
