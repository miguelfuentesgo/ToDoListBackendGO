package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gitlab.com/miguelit0/toDoApp/handlers"
	"gitlab.com/miguelit0/toDoApp/server"
)

func main() {
	fmt.Println("Hello mom, im doing backend")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	JWTSECRET := os.Getenv("JWT_SECRET")

	DATABASEURL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWTSECRET,
		Port:        PORT,
		DataBaseUrl: DATABASEURL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler()).Methods(http.MethodGet)
}
