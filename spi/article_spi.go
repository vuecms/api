package spi

import (
	"api/model"
)

type ArticleSpi struct {
	App *Application
}

func NewArticleSpi(app *Application) *ArticleSpi {
	var spi ArticleSpi
	spi.App = app
	return &spi
}

func (q *ArticleSpi) ListArticlesPaginate(page int, limit int) *Paginator {
	var article []model.Article
	db := model.NewArticleQuerySet(q.App.DB).DB()
	paginator := Paging(&Param{
		DB:    db,
		Page:  page,
		Limit: limit,
		//OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &article)

	return paginator
}
