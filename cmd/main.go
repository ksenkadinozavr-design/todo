package main

import (
	"github.com/ksenkadinozavr-design/todo"
	"github.com/ksenkadinozavr-design/todo/pkg/handler"
	"log"
)

func main() {
	srv := new(todo_list.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
