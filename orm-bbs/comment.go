package orm_bbs

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model

	PostID int
	Post   Post

	Body string `gorm:"VARCHAR"`
}
