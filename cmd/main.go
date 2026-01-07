package main

import (
	"github.com/ksenkadinozavr-design/todo"
	"github.com/ksenkadinozavr-design/todo/pkg/handler"
	"github.com/ksenkadinozavr-design/todo/pkg/repository"
	"github.com/ksenkadinozavr-design/todo/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_list.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
