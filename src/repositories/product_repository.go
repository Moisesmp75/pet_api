package repositories

import (
	"pet_api/src/database"
	"pet_api/src/models"

	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// var (
// 	validate = validator.New()
// )

func GetAllProduct(c *fiber.Ctx) (*[]models.Product, error) {
	var products []models.Product
	database.DB.Find(&products)
	return &products, nil
}

func GetProduct(c *fiber.Ctx) (*models.Product, error) {
	id := c.Params("id")

	var product models.Product
	result := database.DB.Find(&product, id)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &product, nil
}

func CreateProduct(c *fiber.Ctx) (*models.Product, error) {
	newProduct := new(models.Product)

	if err := c.BodyParser(newProduct); err != nil {
		return nil, err
	}

	// if err := validate.Struct(newProduct); err != nil {
	// 	return nil, err
	// }

	if err := database.DB.Create(&newProduct).Error; err != nil {
		return nil, err
	}

	return newProduct, nil
}
