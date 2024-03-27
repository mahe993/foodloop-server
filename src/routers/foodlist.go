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

	r.Route("/{userID}", func(r chi.Router) {
		r.Use(Foodlist.GetUserID)
		r.Get("/", services.Foodlist.GetAllForUser)
		r.Post("/", services.Foodlist.CreateFoodlist)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(Foodlist.GetUserIDAndFoodlistID)
			r.Get("/", services.Foodlist.GetFoodlist)
			r.Put("/status", services.Foodlist.UpdateFoodlistStatus)
			r.Put("/index", services.Foodlist.UpdateFoodlistIndex)
			r.Delete("/", services.Foodlist.DeleteFoodlist)
		})
	})

	return r
}

func (*FoodlistRouter) GetUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID := chi.URLParam(r, "userID")

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (*FoodlistRouter) GetUserIDAndFoodlistID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		foodlistID := chi.URLParam(r, "id")
		userID := chi.URLParam(r, "userID")
		if foodlistID == "" {
			render.Respond(w, r, errors.New("foodlist ID is required"))
			return
		}

		ctx := context.WithValue(r.Context(), "foodlistID", foodlistID)
		ctx = context.WithValue(ctx, "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
