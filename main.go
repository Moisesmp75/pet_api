package main

import (
	"fmt"
	"log"
	"os"
	_ "pet_api/docs"
	"pet_api/src/controllers"
	"pet_api/src/database"
	"pet_api/src/database/migration"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title						HairyPets API
// @version					1.0
// @description				This is a HairyPets API swagger
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.email				fiber@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:3000
// @BasePath					/api/v1
// @securityDefinitions.apikey	ApiKeyAuth
// @In							header
// @Name						Authorization
// @Description				Enter Bearer {token}
func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Could no load .env file")
	}

	database.InitDatabase()
	migration.RunMigration()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

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
