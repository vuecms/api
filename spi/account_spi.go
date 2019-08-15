package spi

import (
	"api/model"
)

type AccountSpi struct {
	App *Application
}

func NewAccountSpi(app *Application) *AccountSpi {
	var spi AccountSpi
	spi.App = app
	return &spi
}

func (q *AccountSpi) ListAccountsPaginate(page int, limit int) *Paginator {
	var account []model.Account
	db := model.NewAccountQuerySet(q.App.DB).DB()
	paginator := Paging(&Param{
		DB:    db,
		Page:  page,
		Limit: limit,
		//OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &account)

	return paginator
}
