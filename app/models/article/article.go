package article

import (
	"goblog/app/models"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	Title      string `gorm:"type:text;not null" valid:"title"`
	Body       string `gorm:"type:text;not null" valid:"body"`
	UserID     uint64
	User       user.User         `gorm:"foreignKey:UserID"`
	Category   category.Category `gorm:"foreignKey:CategoryID"`
	CategoryID uint64            `valid:"type"`
}

/**
type Object struct {
	...
}
// Object 的方法
func (obj *Object) method() {
	...
}

// 只是一个函数
func function() {
	...
}

调用的对比
// 调用方法：
o := new(Object)
o.method()

// 调用函数
function()
*/

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(article.ID), 10))
}

// CreatedAtDate 创建日期
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}
