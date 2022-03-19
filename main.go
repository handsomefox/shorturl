package main

import (
	"log"
	"shorturl/server"
)

func main() {
	// Create the server at the required address, run and log the errors if any happen.
	router := server.Server{Address: "localhost:3000"}
	log.Fatal(router.Run())
}
