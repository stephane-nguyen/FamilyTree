package handler

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

type AuthHandler struct {
	// No need for storage related to Oauth
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) AuthCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendString(user.Email)
}

func (a *AuthHandler) LogoutFromProvider(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Fatal(err)
	}

	return c.SendString("logout")
}