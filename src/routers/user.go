package routers

import (
	"context"
	"errors"
	"foodloop/src/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserRouter struct{}

var User UserRouter

func (*UserRouter) Router() chi.Router {

	r := chi.NewRouter()

	r.Get("/", services.User.GetAll)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(User.IDCtx)
		r.Get("/", services.User.GetUser)
	})

	return r
}

func (*UserRouter) IDCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID := chi.URLParam(r, "id")
		if userID == "" {
			render.Respond(w, r, errors.New("user ID is required"))
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
