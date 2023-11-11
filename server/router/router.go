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
    personsV1.Get("", personHandlers.GetAllPersons)
    personsV1.Get("/:id", personHandlers.GetPerson)
    personsV1.Get("/country/:country", personHandlers.GetpersonsByCountry)
    personsV1.Post("", personHandlers.CreatePerson)
    personsV1.Put("/:id", personHandlers.UpdatePerson)
    personsV1.Delete("/:id", personHandlers.DeletePerson)
	
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("[SERVER PORT] Starting fiber server on port: " + env.SERVER_PORT)
			go app.Listen(":" + env.SERVER_PORT)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
