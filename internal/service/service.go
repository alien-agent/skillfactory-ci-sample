package service

import (
	"fmt"
	"net/http"
)

type HelloServicer interface {
	Hello() string
}

type HelloService struct{}

func (s HelloService) Hello() string {
	return "Hello, World!"
}

func HelloWorldHandler(service HelloServicer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, service.Hello())
	}
}
