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
	"fmt"
	"foodloop/src/database"
	"foodloop/src/routers"
	"log"
	"net/http"
	"os"
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
		log.Print("No .env file found: ", err)
	}
	database.InitDB()
	defer database.CloseDB()

	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// TODO: change example.com to FE domain once deployed
		AllowedOrigins:   []string{"http://localhost:5173", "https://foodloop-6zox.onrender.com"},
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
	r.Mount(basePath+"/foodlist", routers.Foodlist.Router())
	log.Printf("Starting server on port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
