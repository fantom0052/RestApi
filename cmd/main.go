package main

import (
	"github.com/fantom0052/RestApi.git"
	"github.com/fantom0052/RestApi.git/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
