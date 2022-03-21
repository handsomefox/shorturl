package main

import (
	"fmt"
	"log"
	"os"
	"shorturl/internal/server"
)

func main() {
	s := server.ShortURLServer{Address: ":8080", DBKey: os.Getenv("MONGO_KEY")}
	s.Init()
	fmt.Println("Started the server at port 8080")
	log.Fatal(s.Run())
}
