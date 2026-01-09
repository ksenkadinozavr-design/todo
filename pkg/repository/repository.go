package repository

import (
	"github.com/jmoiron/sqlx"
	todo_list "github.com/ksenkadinozavr-design/todo"
	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(user todo_list.User) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
