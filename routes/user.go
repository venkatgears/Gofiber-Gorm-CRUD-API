package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/database"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/models"
)

type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseuser(userModel models.User) UserSerializer {
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseuser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []UserSerializer{}
	for _, user := range users {
		responseuser := CreateResponseuser(user)
		responseUsers = append(responseUsers, responseuser)

	}

	return (c.Status(200).JSON(responseUsers))
}

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("int is an insteger ??")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseuser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("int is an insteger ??")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	var updatedata UpdateUser

	if err := c.BodyParser(&updatedata); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updatedata.FirstName
	user.LastName = updatedata.LastName

	database.Database.Db.Save(&user)

	responseuser := CreateResponseuser(user)
	return c.Status(200).JSON(responseuser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("please ensure that id is INteger")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("successfully deleted user")
}
