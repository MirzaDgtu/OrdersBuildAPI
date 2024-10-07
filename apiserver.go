package main

import (
	"log"
	"ordersbuild/internal/apiserver"
)

func main() {
	log.Println("Main log....")
	log.Fatal(apiserver.RunAPI("127.0.0.1:8090"))
}
