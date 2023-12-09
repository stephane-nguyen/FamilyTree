package handler

import (
	"strconv"
    "github.com/gofiber/fiber/v2"
)

type basicResponse struct {
	Success bool `json:"success"`
}


func GetId(c *fiber.Ctx) (int, error) {
	id := c.Params("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return intID, nil
}

func GetRequestBody[T any](c *fiber.Ctx) (T, error) {
    var reqBody T

    err := c.BodyParser(&reqBody)
    if err != nil {
        // Return an empty instance of T and the error
        return *new(T), err
    }

    return reqBody, nil
}