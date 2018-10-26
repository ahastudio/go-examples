package ormbbs

import (
	"fmt"
	"log"

	"github.com/ahastudio/go-examples/orm-bbs/db"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ExampleComment() {
	db, err := db.Open()
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	db.AutoMigrate(&Post{}, &Comment{})

	post := &Post{
		Title: "New Post",
		Body:  "Hello, world!",
	}

	post.Comments = append(post.Comments, Comment{Body: "Hi!"})
	post.Comments = append(post.Comments, Comment{Body: "Nice!"})

	db.Create(post)

	found := Post{}
	db.First(&found, post.ID).Related(&found.Comments)

	fmt.Println(len(found.Comments))
	for _, comment := range found.Comments {
		fmt.Println(comment.Body)
	}

	// Output:
	// 2
	// Hi!
	// Nice!
}
