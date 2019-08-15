package protocol

import (
	"api/model"
	"github.com/gin-gonic/gin"
)

func NewError(ctx *gin.Context, status int, err error) {
	er := Error{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

//for document

type AreaPaginator struct {
	Paginator
	Records []model.Area `json:"records"`
}

type QuestionLibraryPaginator struct {
	Paginator
	Records []model.QuestionLibrary `json:"records"`
}

type AccountPaginator struct {
	Paginator
	Records []model.Account `json:"records"`
}

type ArticlePaginator struct {
	Paginator
	Records []model.Article `json:"records"`
}