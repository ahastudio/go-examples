package ormbbs

import (
	"github.com/jinzhu/gorm"
)

// Comment model.
type Comment struct {
	gorm.Model

	PostID int
	Post   Post

	Body string `gorm:"VARCHAR"`
}
