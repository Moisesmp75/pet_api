package main

import (
	"fmt"
	"log"
	"os"
	"pet_api/src/controllers"
	"pet_api/src/database"
	"pet_api/src/database/migration"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No se pudo cargar el archivo .env")
	}

	database.InitDatabase()
	migration.RunMigration()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
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

	port := os.Getenv("PORT")
	err := app.Listen(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err.Error())
	}
}
