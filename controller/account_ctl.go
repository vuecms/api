package controller

import (
	"github.com/gin-gonic/gin"
)


// ListAccounts godoc
// @Summary List Account
// @Description get Account libs
// @Tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} protocol.AccountPaginator
// @Failure 400 {object} protocol.Error
// @Failure 404 {object} protocol.Error
// @Failure 500 {object} protocol.Error
// @Router /account [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	paginator := c.App.AccountSpi.ListAccountsPaginate(1,3)
	ctx.JSON(200, paginator)
}
