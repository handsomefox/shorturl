package main

import (
	"fmt"
	"log"
	"os"
	"shorturl/pkg/server"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = ":8080"
	} else {
		port = ":" + port
	}

	key := os.Getenv("MONGO_KEY")

	s := server.New(port, key)
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Started server at port: %s\n", port)
	log.Fatal(s.Run())
}
