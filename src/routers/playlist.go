package routers

import (
	"context"
	"errors"
	"foodloop/src/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type FoodlistRouter struct{}

var Foodlist FoodlistRouter

func (*FoodlistRouter) Router() chi.Router {

	r := chi.NewRouter()

	r.Get("/", services.Foodlist.GetAll)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(Foodlist.IDCtx)
		r.Get("/", services.Foodlist.GetFoodlist)
	})

	return r
}

func (*FoodlistRouter) IDCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		foodlistID := chi.URLParam(r, "id")
		if foodlistID == "" {
			render.Respond(w, r, errors.New("foodlist ID is required"))
			return
		}

		ctx := context.WithValue(r.Context(), "foodlist", foodlistID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
