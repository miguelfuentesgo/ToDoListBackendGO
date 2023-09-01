package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"

	"gitlab.com/miguelit0/toDoApp/router"
)

func main() {
	fmt.Println("Hello mom, im doing backend")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	router := router.NewRouter()

	handler := router.Handler

	log.Printf("Listening on port %v", PORT)

	if err := fasthttp.ListenAndServe("localhost:"+PORT, handler); err != nil {
		log.Fatal(err.Error())
	}

}
