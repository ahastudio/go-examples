package orm_bbs

import (
	"fmt"
	"log"

	"github.com/ahastudio/go-examples/orm-bbs/db"
)

func ExamplePost() {
	db, err := db.Open()
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	db.AutoMigrate(&Post{})

	post := &Post{
		Title: "New Post",
		Body:  "Hello, world!",
	}

	db.Create(post)

	found := Post{}
	db.First(&found, post.ID)

	fmt.Println(post.Title)
	fmt.Println(found.Title)

	// Output:
	// New Post
	// New Post
}
