package model

import (
	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in question.go -out question_query.go

// gen:qs
type QuestionLibrary struct {
	gorm.Model
	Subject       int
	QuestionType  int
	TechType      int
	SecurityLevel int
	Status        int
	Difficulty    int
	InspectPoint  string
}

// gen:qs
type Question struct {
	gorm.Model
	SortIndex int
	Type      int
	Content   int
	Analysis  string
}
