package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null" valid:"title"`
	Body  string `gorm:"type:longtext;not null" valid:"body"`
}

func (a Article) Link() string {
	return route.RouteName2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
