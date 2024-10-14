package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevitoDbug/golangJWTAuthTemplate/routes"
	"github.com/DevitoDbug/golangJWTAuthTemplate/utils"
)

func main() {
	port := ":8080"

	var router http.HandlerFunc = routes.Router
	fmt.Printf("Starting server on port %v....\n", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		customError := &utils.Error{
			Context: "main@main",
			Info:    fmt.Sprintf("failed to start server at port %v", port),
		}

		log.Fatal(customError.Error())
	}
}
