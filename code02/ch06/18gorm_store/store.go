package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id        int
	Content   string
	Author    string    `gorm:"not null"`
	Comments  []Comment `gorm:"foreignKey:PostId"`
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `gorm:"not null"`
	PostId    int    `gorm:"not null"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	dsn := "user=gwp dbname=gwp password=gwp sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	// 新規投稿の作成
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	// コメントの追加
	comment := Comment{Content: "いい投稿だね！", Author: "Joe", PostId: post.Id}
	Db.Create(&comment)

	// 投稿とそのコメントを取得
	var readPost Post
	err := Db.Preload("Comments").Where("author = ?", "Sau Sheong").First(&readPost).Error
	if err != nil {
		panic(err)
	}

	// コメントの確認
	for _, c := range readPost.Comments {
		fmt.Println(c)
	}
}
