package spi

import (
	"api/model"
	"github.com/jinzhu/gorm"
	"log"
)

const DefaultConfigFile = "config/app.toml"

type Application struct {
	Config      *AppConfig
	DB          *gorm.DB
	AreaSpi     *AreaSpi
	QuestionSpi *QuestionSpi
}

func NewInstance(file string) (*Application, error) {
	var app Application
	config, err := Load(file)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(config.Database.Driver, config.Database.Url)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	app.DB = db

	db.AutoMigrate(
		&model.Area{},
		&model.Organization{},
		&model.QuestionLibrary{},
	)

	app.Config = config
	app.QuestionSpi = NewQuestionSpi(&app)
	app.AreaSpi = NewAreaSpi(&app)

	return &app, err
}

func DefaultInstance() (*Application, error) {
	return NewInstance(DefaultConfigFile)
}

func (app *Application) Destroy() {
	defer func() {
		err := app.DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
