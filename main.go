package main

import (
	"github.com/sergey4qb/mf1-test/application"
	"log"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	app, err := application.New()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
