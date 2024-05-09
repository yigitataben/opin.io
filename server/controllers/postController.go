package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yigitataben/go-jwt/initializers"
	"github.com/yigitataben/go-jwt/models"
	"gorm.io/gorm"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Content      string
		CategoryName string
		CategoryCode uint
		UserID       uint
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	post := models.Post{
		Content:      body.Content,
		CategoryName: body.CategoryName,
		CategoryCode: body.CategoryCode,
		UserID:       body.UserID,
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Order("created_at desc").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve posts",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Post not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Post not found",
		})
		return
	}
	var newPostData models.Post
	if err := c.Bind(&newPostData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	post.Content = newPostData.Content
	post.CategoryName = newPostData.CategoryName
	post.CategoryCode = newPostData.CategoryCode
	post.UserID = newPostData.UserID
	result = initializers.DB.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	result = initializers.DB.Unscoped().Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
