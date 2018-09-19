package orm_bbs

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model

	Title string `gorm:"Size:100"`
	Body  string `gorm:"VARCHAR"`

	Comments []Comment
}
