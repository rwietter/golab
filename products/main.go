package main

import (
	"fmt"
	"log"

	"api/app"
)

func main() {
	srv := app.CreateServer()

	fmt.Println("Server running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
