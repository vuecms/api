package controller

import (
	"gopkg.in/go-playground/validator.v9"
)

type AddQuestionParam struct {
	Name        string `form:"name" json:"name" binding:"required"`
	NestedParam NestedParam
}

type NestedParam struct {
	Nested1 string `json:"nested1"`
	Nested2 int    `form:"n2" json:"nested2" binding:"required,gt=0"`
}

//to use this function
//validate.RegisterValidation("is-unique", ValidateUnique)
func ValidateUnique(fl validator.FieldLevel) bool {
	return true
}
