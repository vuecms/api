package controller

import (
	"github.com/gin-gonic/gin"
)


// ListArticle godoc
// @Summary List Article
// @Description get Articles
// @Tags article
// @Accept  json
// @Produce  json
// @Success 200 {object} protocol.ArticlePaginator
// @Failure 400 {object} protocol.Error
// @Failure 404 {object} protocol.Error
// @Failure 500 {object} protocol.Error
// @Router /article [get]
func (c *Controller) ListArticles(ctx *gin.Context) {
	paginator := c.App.ArticleSpi.ListArticlesPaginate(1,3)
	ctx.JSON(200, paginator)
}
