package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	todo_list "github.com/ksenkadinozavr-design/todo"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo_list.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id")
	var row = r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
