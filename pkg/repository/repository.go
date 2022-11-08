package repository

import "github.com/jmoiron/sqlx"

type License struct {
}

type Repository struct {
	License
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
