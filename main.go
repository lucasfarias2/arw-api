package main

import (
	"arw-api/db"
	"arw-api/handlers"
	"arw-api/middleware"
	"arw-api/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	_ = utils.LoadEnv(".env")

	db.ConnectDatabase()

	router.Use(middleware.ValidateToken())

	router.With(middleware.RequireUser).Get("/api/v1/user", handlers.GetUser)
	router.Post("/api/v1/register", handlers.Register)
	router.Post("/api/v1/login", handlers.Login)

	router.NotFound(handlers.NotFound)

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Error starting the server", err)
		return
	}
}
