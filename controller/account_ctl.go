package controller

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) ListAccounts(ctx *gin.Context) {
	paginator := c.App.AccountSpi.ListAccountsPaginate(1,3)
	ctx.JSON(200, paginator)
}
