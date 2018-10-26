package ormbbs

import (
	"github.com/jinzhu/gorm"
)

// Post model.
type Post struct {
	gorm.Model

	Title string `gorm:"Size:100"`
	Body  string `gorm:"VARCHAR"`

	Comments []Comment
}
