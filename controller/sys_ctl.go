package controller

import (
	"api/protocol"
	"github.com/gin-gonic/gin"
)

// CreateArea godoc
// @Summary Create an area
// @Description Create an area
// @Tags area
// @Accept  json
// @Produce  json
// @Param body body protocol.AreaCreationParam true "Add Area"
// @Success 201 {object} protocol.Success
// @Failure 400 {object} protocol.Error
// @Failure 404 {object} protocol.Error
// @Failure 500 {object} protocol.Error
// @Router /areas [put]
func (c *Controller) CreateArea(ctx *gin.Context) {
	//b := binding.Default(ctx.Request.Method, ctx.ContentType())
	pram := &protocol.AreaCreationParam{}
	//err := ctx.ShouldBind(s)
	err := ctx.ShouldBindJSON(pram)
	if err != nil {
		ctx.JSON(400, NewValidatorError(err))
		return
	}
	err = c.App.AreaSpi.CreateArea(pram)
	if err != nil {
		return
	}
	ctx.JSON(201, nil)
}

// Area godoc
// @Summary List Area
// @Description get areas
// @Tags area
// @Accept  json
// @Produce  json
// @Success 200 {array} protocol.AreaPaginator
// @Failure 400 {object} protocol.Error
// @Failure 404 {object} protocol.Error
// @Failure 500 {object} protocol.Error
// @Router /areas [get]
func (c *Controller) ListAreas(ctx *gin.Context) {
	paginator := c.App.AreaSpi.ListArea(1, 10)
	ctx.JSON(200, paginator)
}
