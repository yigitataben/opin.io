package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
	"time"
)

type PostResponse struct {
	ID           uint      `json:"post_id"`
	CreatedAt    time.Time `json:"created_at"`
	Content      string    `json:"content"`
	UserID       uint      `json:"user_id"`
	UserName     string    `json:"user_name"`
	CategoryID   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

func CreateResponsePost(postModel models.Post) PostResponse {
	return PostResponse{
		ID:           postModel.PostID,
		CreatedAt:    postModel.CreatedAt,
		Content:      postModel.Content,
		UserID:       postModel.UserID,
		UserName:     postModel.UserName,
		CategoryID:   postModel.CategoryID,
		CategoryName: postModel.CategoryName,
	}
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&post)
	responsePost := CreateResponsePost(post)

	return c.Status(200).JSON(responsePost)
}

func GetAllPosts(c *fiber.Ctx) error {
	posts := []models.Post{}

	database.Database.DB.Find(&posts)
	responsePosts := []PostResponse{}

	for _, post := range posts {
		responsePost := CreateResponsePost(post)
		responsePosts = append(responsePosts, responsePost)
	}

	return c.Status(200).JSON(responsePosts)
}

func FindPost(id int, post *models.Post) error {
	database.Database.DB.Find(&post, "post_id = ?", id)
	if post.PostID == 0 {
		return errors.New("Post does not exist.")
	}
	return nil
}

func GetPostByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var post models.Post

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer.")
	}
	if err := FindPost(id, &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responsePost := CreateResponsePost(post)

	return c.Status(200).JSON(responsePost)
}

func UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var post models.Post

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer.")
	}
	if err := FindPost(id, &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdatePost struct {
		Content string `json:"content"`
	}
	var updateData UpdatePost

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	post.Content = updateData.Content

	database.Database.DB.Save(&post)

	responsePost := CreateResponsePost(post)
	return c.Status(200).JSON(responsePost)
}

func DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var post models.Post

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer.")
	}
	if err := FindPost(id, &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.DB.Delete(&post).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Post successfully deleted.")
}
