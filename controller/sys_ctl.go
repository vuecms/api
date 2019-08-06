package controller

import (
	"api/protocol"
	"github.com/gin-gonic/gin"
)


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

func (c *Controller) ListAreas(ctx *gin.Context) {
	paginator := c.App.AreaSpi.ListArea(1, 10)
	ctx.JSON(200, paginator)
}
