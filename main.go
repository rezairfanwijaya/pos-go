package main

import (
	"log"
	"pos/database"
)

func main() {
	connection, err := database.NewConnection(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(connection)
}
