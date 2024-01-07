/*
Package services consumes the requests and provides the business logic pertaining to the request.
*/
package services

import (
	"foodloop/src/database"
	"foodloop/src/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type ExampleService struct{}

var Example ExampleService

// Index is the handler for GET /api/v1/example
func (*ExampleService) Index(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "Hello World")
}

// JSON is the handler for GET /api/v1/example/json
func (*ExampleService) JSON(w http.ResponseWriter, r *http.Request) {
	// example JSON
	exampleStruct := struct {
		Message string `json:"message"`
	}{
		Message: "Hello World",
	}

	// Notice that we still receive the JSON even though status is set randomly to 202
	render.Status(r, http.StatusAccepted)
	render.JSON(w, r, exampleStruct)
}

// ReceiveJSON is the handler for POST /api/v1/example/json
func (*ExampleService) ReceiveJSON(w http.ResponseWriter, r *http.Request) {
	exampleDataToReceive := &models.ExampleModel{}

	if err := render.DecodeJSON(r.Body, exampleDataToReceive); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	log.Println(exampleDataToReceive)
}

// GetID is the handler for GET /api/v1/example/{id}
func (*ExampleService) GetID(w http.ResponseWriter, r *http.Request) {
	// Retrieve context value and type assert it to string
	// Notice that we can potentially pass a struct as the context value adn type assert it to the struct
	// This is useful if we have a bunch of routes that require the same context value
	id := r.Context().Value("exampleId").(string)
	render.Respond(w, r, id)
}

// UpdateID is the handler for PUT /api/v1/example/{id}
func (*ExampleService) UpdateID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("exampleId").(string)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// This should ideally be a custom error object
		errStruct := struct {
			Error string `json:"error"`
			Code  int    `json:"code"`
			Hint  string `json:"hint"`
		}{
			Error: err.Error(),
			// This should ideally be the appropriate API error code
			// https://deliveryhero.udemy.com/course/rest-api/learn/lecture/21525684#overview
			Code: 1000,
			Hint: "Please check the ID, it should be an integer",
		}
		render.Status(r, http.StatusBadRequest)
		// We can also use render.Render() to render a custom error object
		// As long as the error object implements the render.Renderer interface
		render.JSON(w, r, errStruct)
		return
	}

	newIDStruct := struct {
		OldID int `json:"oldId"`
		NewID int `json:"newId"`
	}{
		OldID: idInt,
		NewID: idInt + 1,
	}

	render.JSON(w, r, newIDStruct)
}

// GetFood is the handler for GET /api/v1/example/food
func (*ExampleService) GetFood(w http.ResponseWriter, r *http.Request) {
	rows, err := database.GetDB().Query("SELECT * FROM food")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var food struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var res []struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	for rows.Next() {
		err := rows.Scan(&food.Id, &food.Name, &food.Description)
		if err != nil {
			panic(err)
		}
		res = append(res, food)
	}

	render.JSON(w, r, res)
}
