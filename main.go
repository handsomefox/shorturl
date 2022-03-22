package main

import (
	"fmt"
	"log"
	"os"
	"shorturl/internal/server"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = ":8080"
	} else {
		port = ":" + port
	}

	key := os.Getenv("MONGO_KEY")

	s := server.ShortURLServer{Address: port, DBKey: key}
	err := s.Init()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Started server at port: %s\n", port)
	log.Fatal(s.Run())
}
