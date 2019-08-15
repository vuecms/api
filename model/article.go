package model

import (
	"github.com/jinzhu/gorm"
	"html/template"
)

//go:generate goqueryset -in article.go -out article_query.go

// gen:qs
type Article struct {
	gorm.Model
	PostType    string // 文章类型
	Subject     string // 标题
	Picture     string // 图片
	Description string // 描述
	Body        template.HTML //内容
 	UserID      int // 用户ID
	NickName    string // 用户昵称
	Origin      string // 用户名
	AddDate     string // 创建时间
	Visited     int   // 访问次数
	Status		int    // 状态
}