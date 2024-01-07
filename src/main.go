/*
Package main is the entry point for foodloop server.
Middlewares are defined within the scope of this package.
Routes are directed to package router for further parsing.

Primary module:

  - main.go

Other modules:
  - N.A.
*/
package main

import (
	"foodloop/src/database"
	"foodloop/src/routers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

const (
	basePath = "/api/v1"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()
	defer database.CloseDB()

	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// TODO: change example.com to FE domain once deployed
		AllowedOrigins:   []string{"http://www.example.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}))

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// Sets a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(10 * time.Second))

	// mount routes
	r.Mount(basePath+"/example", routers.Example.Router())
	r.Mount(basePath+"/user", routers.User.Router())

	log.Println("Starting server on port 1111")
	log.Fatal(http.ListenAndServe(":1111", r))
}
