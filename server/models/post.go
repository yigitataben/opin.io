package models

import (
	"github.com/jinzhu/gorm"
	"github.com/yigitataben/opin.io/config"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Category string `gorm:"json:category"`
	Title    string `gorm:"json:title"`
	Post     string `gorm:"json:post"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (p *Post) CreatePost() *Post {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllPosts() []Post {
	var Posts []Post
	db.Find(Posts)
	return Posts
}

func GetPostByID(ID int64) (*Post, gorm.DB) {
	var getPost Post
	db := db.Where("ID=?", ID).Find(&getPost)
	return &getPost, *db
}

func DeletePost(ID int64) Post {
	var post Post
	db.Where("ID=?", ID).Delete(post)
	return post
}
