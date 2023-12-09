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

func NewFiberServer(lc fx.Lifecycle, memberHandlers *handler.MemberHandler) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	env := config.LoadEnv()

	v1 := app.Group("/v1") 
    membersV1 := v1.Group("/members") 
    membersV1.Get("", memberHandlers.GetAllMembers)
    membersV1.Get("/:id", memberHandlers.GetMember)
    membersV1.Get("/country/:country", memberHandlers.GetMembersByCountry)
    membersV1.Post("", memberHandlers.CreateMember)
    membersV1.Put("/:id", memberHandlers.UpdateMember)
    membersV1.Delete("/:id", memberHandlers.DeleteMember)
	
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
