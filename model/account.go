package model


import (
	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in account.go -out account_query.go

// gen:qs
type Account struct {
	gorm.Model
	UserName      string
	Age   int
	Pass  string
}
