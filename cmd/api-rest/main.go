package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	database "gitlab.com/miguelit0/toDoApp/database"
	repository "gitlab.com/miguelit0/toDoApp/repository"
	"gitlab.com/miguelit0/toDoApp/router"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Saving in variables from .env
	PORT := os.Getenv("PORT")
	URL := os.Getenv("DATABASE_URL")

	//struct database with methods that execute queries and DML
	repo, err := database.NewPostgresRepository(URL)

	if err != nil {
		log.Fatal(err)
	}

	// Assign postgres repository as repository. It complies the interface repository
	repository.NewRepository(repo)

	routerTasks := router.NewRouter()

	handler := routerTasks.Handler

	log.Printf("Listening on port %v", PORT)

	if err := fasthttp.ListenAndServe("localhost:"+PORT, handler); err != nil {
		log.Fatal(err.Error())
	}

}
