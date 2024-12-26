package main

import (
	"fmt"
	"net/http"

	"skillfactory-ci-sample/internal/service"
)

func main() {
	svc := service.HelloService{}
	http.HandleFunc("/", service.HelloWorldHandler(svc))

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
