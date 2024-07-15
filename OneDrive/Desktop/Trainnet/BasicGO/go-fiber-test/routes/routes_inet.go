package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v3 := api.Group("/v3")
	v2.Get("/", c.HelloTestV2)
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.ParamsTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidationTest)

	//5.1
	v1.Get("/fact/:n", c.FactorialTest)
	//5.2
	v3.Post("/:name", c.QueryTestId)
	//6
	v3.Post("/register", c.RegisterTest)
}
