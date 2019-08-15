package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	//*gorm.Model
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

func (qs AreaQuerySet) DB() *gorm.DB {
	return qs.db
}

func (qs AccountQuerySet) DB() *gorm.DB {
	return qs.db
}

func (qs QuestionLibraryQuerySet) DB() *gorm.DB {
	return qs.db
}
