package repository

import (
	"github.com/dgolov/LicenseServer"
	"github.com/jmoiron/sqlx"
)

type License interface {
	GetLicenseByUuid(uuid string) (LicenseServer.License, error)
}

type Repository struct {
	License
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		License: NewCheckPostgres(db),
	}
}
