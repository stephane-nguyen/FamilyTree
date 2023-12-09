package handler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/stephane-nguyen/FamilyTree/server/storage"
)

type UserHandler struct{
	Storage *storage.UserStorage
}

type fetchUsersResponse struct {
	Users []storage.User_DB `json:"users"`
}

type fetchOneUserResponse struct {
	User storage.User_DB `json:"user"`
}

type createUserRequest struct {
	Username   string    `json:"username" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Email      string    `json:"email" validate:"required"`
}

type createUserResponse struct {
	Id 			int		 `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
}

func NewUserHandler(storage *storage.UserStorage) *UserHandler {
	return &UserHandler{Storage: storage}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	Users, err := h.Storage.GetAllUsers()
	if err != nil {
		return err
	}

	response := fetchUsersResponse{
		Users: Users,
	}
	return c.JSON(response)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }

	User, err := h.Storage.GetUserById(id)
	if err != nil {
		return err
	}

	response := fetchOneUserResponse{User: User}
	return c.JSON(response)
}


func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	requestBody, err := GetRequestBody[createUserRequest](c)
    if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }
	
	id, err := h.Storage.CreateNewUser(storage.NewUserInput{
		Username: requestBody.Username,
		Password: requestBody.Password,
		Email: requestBody.Email,
	})
	if err != nil {
		return err
	}

	resp := createUserResponse{Id: id, Username: requestBody.Username, Email: requestBody.Email}

	return c.JSON(resp)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }

	err = h.Storage.UpdateUserById(id)
	if err != nil {
		return err
	}

	resp := basicResponse{
		Success: true,
	}

	return c.JSON(resp)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
		return err
	}

	err = h.Storage.DeleteUserById(id)
	if err != nil {
		return err
	}

	resp := basicResponse{
		Success: true,
	}

	return c.JSON(resp)
}
