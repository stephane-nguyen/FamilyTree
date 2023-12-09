package router

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shareed2k/goth_fiber"
	"github.com/stephane-nguyen/FamilyTree/server/config"
	"github.com/stephane-nguyen/FamilyTree/server/handler"
	"github.com/stephane-nguyen/FamilyTree/server/auth"
)

func NewFiberServer(
	lc fx.Lifecycle, 
	authHandlers *handler.AuthHandler, 
	memberHandlers *handler.MemberHandler, 
	userHandlers *handler.UserHandler,
	) *fiber.App {

	env := config.LoadEnv()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	auth.OAuth2()

	v1 := app.Group("/v1") 
	authV1 := v1.Group("/auth")
	usersV1 := v1.Group("/users")
    membersV1 := v1.Group("/members") 

	authV1.Get("/:provider/callback", authHandlers.AuthCallback)
	authV1.Get("/logout/:provider", authHandlers.LogoutFromProvider)
	authV1.Get("/:provider", goth_fiber.BeginAuthHandler)

	usersV1.Post("", userHandlers.CreateUser)

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
