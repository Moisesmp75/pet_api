package main

import (
	"log"
	"pet_api/src/controllers"
	"pet_api/src/database"
	"pet_api/src/database/migration"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.InitDatabase()
	migration.RunMigration()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	apiRoute := app.Group("/api")
	apiV1 := apiRoute.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		c.Set("Callback-Token", "some-token-here")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	controllers.AddControllers(apiV1)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
