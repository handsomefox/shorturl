package main

import (
	"fmt"
	"log"
	"os"
	"shorturl/pkg/server"
)

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8080"
	} else {
		port = ":" + port
	}
	return port
}

func main() {
	port := getPort()
	key := os.Getenv("MONGO_KEY")
	if len(key) == 0 {
		log.Fatal("no database key was provided")
	}

	s := server.New(port, key)
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Started server at port: %s\n", port)
	log.Fatal(s.Run())
}
