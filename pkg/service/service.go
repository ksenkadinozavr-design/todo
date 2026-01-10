package service

import (
	todo_list "github.com/ksenkadinozavr-design/todo"
	"github.com/ksenkadinozavr-design/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_list.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
