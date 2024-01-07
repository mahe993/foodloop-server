/*
Package router defines routers for each resource and routes requests to the respective services.
*/
package routers

import (
	"context"
	"errors"
	"foodloop/src/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ExampleRouter struct{}

var Example ExampleRouter

// Router returns a chi.Router object that handles all requests to /api/v1/example
func (*ExampleRouter) Router() chi.Router {
	r := chi.NewRouter()

	// GET /api/v1/example
	r.Get("/", services.Example.Index)

	r.Route("/json", func(r chi.Router) {
		// GET /api/v1/example/json
		r.Get("/", services.Example.JSON)
		// POST /api/v1/example/json
		r.Post("/", services.Example.ReceiveJSON)
	})

	// Note that this route is potentially clashing with "/json" route
	// i.e. {id} == "json"
	// To avoid this, we should place static routes above dynamic routes
	// To test this, cut and paste this route snippet above the "/json" route
	// Realize that the "/json" route is no longer accessible
	r.Route("/{id}", func(r chi.Router) {
		// Set context value for all routes in this scope
		r.Use(Example.IDCtx)

		// GET /api/v1/example/{id}
		r.Get("/", services.Example.GetID)
		// PUT /api/v1/example/{id}
		r.Put("/", services.Example.UpdateID)

	})

	r.Get("/food", services.Example.GetFood)

	return r
}

// IDCtx is a middleware that sets the context value for all routes in the scope of the route that calls this middleware
func (*ExampleRouter) IDCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		exampleID := chi.URLParam(r, "id")
		if exampleID == "" {
			render.Respond(w, r, errors.New("example ID is required"))
			return
		}

		ctx := context.WithValue(r.Context(), "exampleId", exampleID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
