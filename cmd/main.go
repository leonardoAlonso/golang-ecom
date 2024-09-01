package main

import (
	"log"

	"github.com/leonardoAlonso/go-ecom/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080", nil) // Pointer to the struct
	if err := server.Run(); err != nil {     // Run on the same memory address
		log.Fatal(err)
	}
}
