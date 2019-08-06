package spi

import (
	"api/model"
	"log"
)

type QuestionSpi struct {
	App *Application
}

func NewQuestionSpi(app *Application) *QuestionSpi {
	var spi QuestionSpi
	spi.App = app
	return &spi
}

func (q *QuestionSpi) ListQuestionLibraries() ([]model.QuestionLibrary, error) {
	var questionLibraries []model.QuestionLibrary
	err := model.NewQuestionLibraryQuerySet(q.App.DB).All(&questionLibraries)
	if err != nil {
		log.Fatal(err)
	}
	return questionLibraries, nil
}

func (q *QuestionSpi) ListQuestionLibsPaginate() *Paginator {
	var questionLibs []model.QuestionLibrary
	db := model.NewQuestionLibraryQuerySet(q.App.DB).OrderAscByCreatedAt().DB()
	paginator := Paging(&Param{
		DB:    db,
		Page:  1,
		Limit: 3,
		//OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &questionLibs)

	return paginator
}
