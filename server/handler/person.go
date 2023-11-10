package handler

import (
    // "strconv"
    "github.com/gofiber/fiber/v2"
    "github.com/stephane-nguyen/FamilyTree/server/storage"
)

type PersonHandler struct{
	Storage *storage.PersonStorage
}

func NewPersonHandler(storage *storage.PersonStorage) *PersonHandler {
  return &PersonHandler{Storage: storage}
}

type fetchPersonsResponse struct {
	Persons []storage.Person_DB `json:"persons"`
}

func (h *PersonHandler) FetchPersons(c *fiber.Ctx) error {
	persons, err := h.Storage.GetAllPersons()
	if err != nil {
		return err
	}

	response := fetchPersonsResponse{
		Persons: persons,
	}
	return c.JSON(response)
}


// type createPersonRequest struct {
// 	Title       string `json:"title" validate:"require,email"`
// 	Description string `json:"description" validate:"require,description"`
// }

// type createPersonResponse struct {
// 	Id int `json:"id"`
// }

// func (h *PersonHandler) CreatePerson(c *fiber.Ctx) error {

// 	// get the request body
// 	var body createPersonRequest
// 	err = c.BodyParser(&body)
// 	if err != nil {
// 		return err
// 	}

// 	id, err := h.Storage.CreateNewPerson(storage.NewPersonInput{
// 		UserId:      session.Id,
// 		Title:       body.Title,
// 		Description: body.Description,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// send the id
// 	resp := createPersonResponse{Id: id}

// 	return c.JSON(resp)
// }

// type fetchOnePersonResponse struct {
// 	Person storage.Person_DB `json:"Person"`
// }

// func (h *PersonHandler) FetchPerson(c *fiber.Ctx) error {
	
// 	id := getId(c)

// 	// get the Person
// 	Person, err := h.Storage.GetOnePerson(aid)
// 	if err != nil {
// 		return err
// 	}

// 	// send
// 	resp := fetchOnePersonResponse{
// 		Person: Person,
// 	}

// 	return c.JSON(resp)
// }


// type basicResponse struct {
// 	Success bool `json:"success"`
// }

// func (h *PersonHandler) CompletePerson(c *fiber.Ctx) error {
// 	id := getId(c)

// 	// complete Person
// 	err = h.Storage.CompletePerson(aid)
// 	if err != nil {
// 		return err
// 	}

// 	// send
// 	resp := basicResponse{
// 		Success: true,
// 	}
// 	return c.JSON(resp)
// }

// func (h *PersonHandler) DeletePerson(c *fiber.Ctx) error {
// 	id := getId(c)

// 	err = h.Storage.DeletePerson(aid)
// 	if err != nil {
// 		return err
// 	}

// 	// send
// 	resp := basicResponse{
// 		Success: true,
// 	}
// 	return c.JSON(resp)
// }

// func getId(c *fiber.Ctx) error {
//     id := c.Params("id")
// 	aid, err := strconv.Atoi(id)
// 	if err != nil {
// 		return err
// 	}
//     return id
// }