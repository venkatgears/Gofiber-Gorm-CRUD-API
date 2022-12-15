package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/database"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/models"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSerializer {
	return ProductSerializer{
		ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func Getproducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)

	responseProducts := []ProductSerializer{}

	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func Findproducts(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id=?", id)

	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("please ensure id in integer")
	}

	if err := Findproducts(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func Updateproduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("please ensure id in integer")
	}

	if err := Findproducts(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	var updateproduct UpdateProduct

	if err := c.BodyParser(&updateproduct); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateproduct.Name
	product.SerialNumber = updateproduct.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("ensure int is an integer")
	}
	if err := Findproducts(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Delete(product)
	return c.Status(200).JSON("delete sucessfull")
}
