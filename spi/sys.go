package spi

import (
	"exam/model"
	"exam/protocol"
	"github.com/jinzhu/copier"
)

type AreaSpi struct {
	App *Application
}

func NewAreaSpi(app *Application) *AreaSpi {
	return &AreaSpi{App: app}
}

func (s *AreaSpi) CreateArea(param *protocol.AreaCreationParam) error {
	var area model.Area
	err := copier.Copy(&area, *param)
	if err != nil {
		return err
	}
	s.App.DB.Save(&area)
	return nil
}

func (s *AreaSpi) ListArea(page int, limit int) *Paginator {
	var areas []model.Area
	db := model.NewAreaQuerySet(s.App.DB).DB()
	return Paging(&Param{
		DB:    db,
		Page:  page,
		Limit: limit,
	}, &areas)
}
