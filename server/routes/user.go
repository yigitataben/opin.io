package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
	"time"
)

type UserResponse struct {
	ID           uint      `json:"id"`
	UserName     string    `json:"user_name"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address"`
	UserPassword string    `json:"user_password"`
	CreatedAt    time.Time `json:"created_at"`
}

func CreateResponseUser(userModel models.User) UserResponse {
	return UserResponse{
		ID:           userModel.UserID,
		UserName:     userModel.UserName,
		FirstName:    userModel.FirstName,
		LastName:     userModel.LastName,
		EmailAddress: userModel.EmailAddress,
		UserPassword: userModel.UserPassword,
		CreatedAt:    userModel.CreatedAt,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.DB.Order("created_at desc").Find(&users)
	responseUsers := make([]UserResponse, len(users))

	for i, user := range users {
		responseUsers[i] = CreateResponseUser(user)
	}

	return c.Status(200).JSON(responseUsers)
}

func FindUser(id int, user *models.User) error {
	result := database.Database.DB.Find(&user, "user_id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("User does not exist.")
	}
	return nil
}

func GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer.")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer.")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		UserName     string `json:"user_name"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		EmailAddress string `json:"email_address"`
		UserPassword string `json:"user_password"`
	}
	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	user.UserName = updateData.UserName
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.EmailAddress = updateData.EmailAddress
	user.UserPassword = updateData.UserPassword

	database.Database.DB.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer.")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.DB.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("User successfully deleted.")
}

func LoginUser(c *fiber.Ctx) error {
	var loginData struct {
		EmailAddress string `json:"email_address"`
		UserPassword string `json:"user_password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid request"})
	}

	var user models.User
	if err := database.Database.DB.Where("email_address = ?", loginData.EmailAddress).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Invalid email or password"})
	}

	if user.UserPassword != loginData.UserPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Invalid email or password"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Login successful"})
}

func LogoutUser(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
