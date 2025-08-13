package repository

import (
	"github.com/Danil-Zlo/todo-list-app"
	"github.com/jmoiron/sqlx"
)

// Интерфесы для работы с бизнес логикой
type Authorization interface {
	CreateUser(user todo.User) (int, error)
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

// конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
