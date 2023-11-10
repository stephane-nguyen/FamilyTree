package router

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/stephane-nguyen/FamilyTree/server/config"
	"github.com/stephane-nguyen/FamilyTree/server/handler"
)

func NewFiberServer(lc fx.Lifecycle, personHandlers *handler.PersonHandler) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	env := config.LoadEnv()

	v1 := app.Group("/v1") 
    personsV1 := v1.Group("/persons") 
    personsV1.Get("/", personHandlers.FetchPersons)
    // persons.Get("/:country", GetpersonsByCountry(driver))
    // persons.Post("/", Create(driver))
    // persons.Update("/:id", Update(driver))
    // persons.Delete("/:id", Delete(driver))
	
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port %d", env.SERVER_PORT)
			go app.Listen(":" + env.SERVER_PORT)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
