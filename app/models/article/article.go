package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null" valid:"title"`
	Body  string `gorm:"type:longtext;not null" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

func (a Article) Link() string {
	return route.RouteName2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}

func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}
