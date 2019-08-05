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
	//var result struct{ Count int }
	//currentField, _, _ := fl.GetStructFieldOK()
	//table := modelTableNameMap[currentField.Type().Name()] // table name
	//value := fl.Field().String()                           // value
	//column := fl.FieldName()                               // column name
	//sql := fmt.Sprintf("select count(*) from %s where %s='%s'", table, column, value)
	//db.PG.Raw(sql).Scan(&result)
	//dup := result.Count > 0
	//return !dup

	//f,_:= fl.GetStructFieldOK()
	//tp := new(f.Type().Name())
	return true
}
