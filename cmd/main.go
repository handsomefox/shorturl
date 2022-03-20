package main

import (
	"log"
	"shorturl/internal/server"
)

func main() {
	// Create the server at the required address, run and log the errors if any happen.
	router := server.ShortURLServer{Address: "localhost:3000"}
	log.Fatal(router.Run())
}
