package handler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/stephane-nguyen/FamilyTree/server/storage"
)

type MemberHandler struct{
	Storage *storage.MemberStorage
}

type fetchMembersResponse struct {
	Members []storage.Member_DB `json:"members"`
}

type fetchOneMemberResponse struct {
	Member storage.Member_DB `json:"member"`
}

type createMemberRequest struct {
	Firstname   string    `json:"firstname" validate:"required"`
	Middlename  string    `json:"middlename"`
	Lastname    string    `json:"lastname" validate:"required"`
	Birthdate   string    `json:"birthdate"`
	Gender      string    `json:"gender" validate:"oneof=Male Female Other"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Photo       []byte    `json:"photo"`
}

type createMemberResponse struct {
	Id 			int		  `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
}

func NewMemberHandler(storage *storage.MemberStorage) *MemberHandler {
	return &MemberHandler{Storage: storage}
}

func (h *MemberHandler) GetAllMembers(c *fiber.Ctx) error {
	members, err := h.Storage.GetAllMembers()
	if err != nil {
		return err
	}

	response := fetchMembersResponse{
		Members: members,
	}
	return c.JSON(response)
}

func (h *MemberHandler) GetMembersByCountry(c *fiber.Ctx) error {
	country := GetCountry(c)
	
	members, err := h.Storage.GetMembersByCountry(country)
	if err != nil {
		return err
	}

	response := fetchMembersResponse{
		Members: members,
	}
	return c.JSON(response)
}

func (h *MemberHandler) GetMember(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }

	Member, err := h.Storage.GetMemberById(id)
	if err != nil {
		return err
	}

	response := fetchOneMemberResponse{Member: Member}
	return c.JSON(response)
}


func (h *MemberHandler) CreateMember(c *fiber.Ctx) error {

	requestBody, err := GetRequestBody[createMemberRequest](c)
    if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }
	
	id, err := h.Storage.CreateNewMember(storage.NewMemberInput{
		Firstname: requestBody.Firstname,
		Middlename: requestBody.Middlename,
		Lastname: requestBody.Lastname,
		Birthdate: requestBody.Birthdate,
		Gender: requestBody.Gender,
		City: requestBody.City,
		Country: requestBody.Country,
	})
	if err != nil {
		return err
	}

	resp := createMemberResponse{Id: id, Firstname: requestBody.Firstname, Lastname: requestBody.Lastname}

	return c.JSON(resp)
}

func (h *MemberHandler) UpdateMember(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }

	err = h.Storage.UpdateMemberById(id)
	if err != nil {
		return err
	}

	resp := basicResponse{
		Success: true,
	}

	return c.JSON(resp)
}

func (h *MemberHandler) DeleteMember(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
		return err
	}

	err = h.Storage.DeleteMemberById(id)
	if err != nil {
		return err
	}

	resp := basicResponse{
		Success: true,
	}

	return c.JSON(resp)
}

func GetCountry(c *fiber.Ctx) string {
    return c.Params("country")
}