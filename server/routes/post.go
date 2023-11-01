package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
)

type Post struct {
	ID           uint   `json:"id"`
	PostCategory string `json:"PostCategory"`
	PostTitle    string `json:"PostTitle"`
	PostBody     string `json:"PostBody"`
}

func CreateResponsePost(postModel models.Post) Post {
	return Post{ID: postModel.ID, PostTitle: postModel.PostTitle, PostBody: postModel.PostBody, PostCategory: postModel.PostCategory}
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&post)
	responseUSer := CreateResponsePost(post)

	return c.Status(200).JSON(responseUSer)
}

func GetAllPosts(c *fiber.Ctx) error {
	posts := []models.Post{}

	database.Database.DB.Find(&posts)
	responsePosts := []Post{}

	for _, post := range posts {
		responsePost := CreateResponsePost(post)
		responsePosts = append(responsePosts, responsePost)
	}

	return c.Status(200).JSON(responsePosts)
}

func FindPost(id int, post *models.Post) error {
	database.Database.DB.Find(&post, "id = ?", id)
	if post.ID == 0 {
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
		PostCategory string `json:"PostCategory"`
		PostTitle    string `json:"PostTitle"`
		PostBody     string `json:"PostBody"`
	}
	var updateData UpdatePost

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	post.PostCategory = updateData.PostCategory
	post.PostTitle = updateData.PostTitle
	post.PostBody = updateData.PostBody

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
