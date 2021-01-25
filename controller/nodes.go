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
	"github.com/gin-gonic/gin"
)

// GetNode godoc
// @Summary Get node by ID
// @Description Get node by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id} [get]
func (c *Controller) GetNode(ctx *gin.Context) {
	// TODO
}

// GetHealth godoc
// @Summary Get node health by ID
// @Description Get node health by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id}/health [get]
func (c *Controller) GetHealth(ctx *gin.Context) {
	// TODO
}

// GetInfo godoc
// @Summary Get node information by ID
// @Description Get node information by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id}/info [get]
func (c *Controller) GetInfo(ctx *gin.Context) {
	// TODO
}

// GetPerf godoc
// @Summary Get node performance by ID
// @Description Get node performance by ID
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes/{id}/perf [get]
func (c *Controller) GetPerf(ctx *gin.Context) {
	// TODO
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
func (c *Controller) QueryNode(ctx *gin.Context) {
	// TODO
}

// AddNode godoc
// @Summary Add node
// @Description Add node
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes [put]
func (c *Controller) AddNode(ctx *gin.Context) {
	// TODO
}

// DelNode godoc
// @Summary Delete node
// @Description Delete node
// @Tags nodes
// @Accept json
// @Produce json
// @Param id path int true "Node ID"
// @Success 200 {object} model.Node
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /nodes [delete]
func (c *Controller) DelNode(ctx *gin.Context) {
	// TODO
}
