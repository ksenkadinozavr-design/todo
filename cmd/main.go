package main

import (
	"github.com/ksenkadinozavr-design/todo"
	"log"
)

func main() {
	srv := new(todo_list.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
