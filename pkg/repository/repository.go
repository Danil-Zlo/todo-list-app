package repository

import "github.com/jmoiron/sqlx"

// Интерфесы для работы с бизнес логикой
type Authorization interface {
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
	return &Repository{}
}
