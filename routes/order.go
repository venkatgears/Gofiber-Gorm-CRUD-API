package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/database"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/models"
)

type OrderSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product`
	CreatedAt time.Time         `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	if err := FindUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := Findproducts(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseuser := CreateResponseuser(user)
	responseproduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseuser, responseproduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Find(&orders)
	responseOrders := []OrderSerializer{}
	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id=?", order.User.ID)
		database.Database.Db.Find(&product, "id=?", order.Product.ID)
		responseorder := CreateResponseOrder(order, CreateResponseuser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseorder)
	}

	return c.Status(200).JSON(responseOrders)

}

func FindOrder(id int ,order *models.Order) error {
	database.Database.Db.Find(&order , "id=?"  ,id)
	if order.ID == 0 {
		errors.New("order not found")
	}
	return nil 
}

func GetOrder(c *fiber.Ctx) error {
	id , err := c.ParamsInt("id")
	var order models.Order
	if err = FindOrder(id ,&order ); err != nil{
		c.Status(400).SendString(err.Error())
	}

	var user models.User
	var product models.Product
	database.Database.Db.Find(order.UserRefer,&user)
	database.Database.Db.Find(order.ProductRefer,&product)
	responseUser := CreateResponseuser(user)
	responseproduct := CreateResponseProduct(product)	
	responseorder := CreateResponseOrder(order, responseUser , responseproduct)

	return c.Status(200).JSON(responseorder)
}