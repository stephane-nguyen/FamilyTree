package handler

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
    "github.com/stephane-nguyen/FamilyTree/server/storage"
)

type PersonHandler struct{
	Storage *storage.PersonStorage
}

type fetchPersonsResponse struct {
	Persons []storage.Person_DB `json:"persons"`
}

type fetchOnePersonResponse struct {
	Person storage.Person_DB `json:"person"`
}

type createPersonRequest struct {
	Firstname   string    `json:"firstname" validate:"required"`
	Middlename  string    `json:"middlename"`
	Lastname    string    `json:"lastname" validate:"required"`
	Birthdate   string    `json:"birthdate"`
	Gender      string    `json:"gender" validate:"oneof=Male Female Other"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Photo       []byte    `json:"photo"`
}

type createPersonResponse struct {
	Id 			int		  `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
}

type basicResponse struct {
	Success bool `json:"success"`
}

func NewPersonHandler(storage *storage.PersonStorage) *PersonHandler {
	return &PersonHandler{Storage: storage}
}

func (h *PersonHandler) GetAllPersons(c *fiber.Ctx) error {
	persons, err := h.Storage.GetAllPersons()
	if err != nil {
		return err
	}

	response := fetchPersonsResponse{
		Persons: persons,
	}
	return c.JSON(response)
}

func (h *PersonHandler) GetpersonsByCountry(c *fiber.Ctx) error {
	country := GetCountry(c)
	
	persons, err := h.Storage.GetpersonsByCountry(country)
	if err != nil {
		return err
	}

	response := fetchPersonsResponse{
		Persons: persons,
	}
	return c.JSON(response)
}

func (h *PersonHandler) GetPerson(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }

	person, err := h.Storage.GetPersonById(id)
	if err != nil {
		return err
	}

	response := fetchOnePersonResponse{Person: person}
	return c.JSON(response)
}


func (h *PersonHandler) CreatePerson(c *fiber.Ctx) error {

	requestBody, err := GetRequestBody(c)
    if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }
	
	id, err := h.Storage.CreateNewPerson(storage.NewPersonInput{
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

	resp := createPersonResponse{Id: id, Firstname: requestBody.Firstname, Lastname: requestBody.Lastname}

	return c.JSON(resp)
}

func (h *PersonHandler) UpdatePerson(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
        c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        return err
    }

	err = h.Storage.UpdatePersonById(id)
	if err != nil {
		return err
	}

	resp := basicResponse{
		Success: true,
	}

	return c.JSON(resp)
}

func (h *PersonHandler) DeletePerson(c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
		return err
	}

	err = h.Storage.DeletePersonById(id)
	if err != nil {
		return err
	}

	resp := basicResponse{
		Success: true,
	}

	return c.JSON(resp)
}

func GetId(c *fiber.Ctx) (int, error) {
	id := c.Params("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return intID, nil
}

func GetCountry(c *fiber.Ctx) string {
    return c.Params("country")
}

func GetRequestBody(c *fiber.Ctx) (createPersonRequest, error) {
    var reqBody createPersonRequest

    err := c.BodyParser(&reqBody); 
	if err != nil {
        return createPersonRequest{}, err
    }

    return reqBody, nil
}