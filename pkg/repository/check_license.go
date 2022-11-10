package repository

import (
	"fmt"
	"github.com/dgolov/LicenseServer"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CheckPostgres struct {
	db *sqlx.DB
}

func NewCheckPostgres(db *sqlx.DB) *CheckPostgres {
	return &CheckPostgres{db: db}
}

func (r *CheckPostgres) GetLicenseByUuid(uuid string) (LicenseServer.License, error) {
	logrus.Printf("[Unit] CheckPostgres - GetLicenseByUuid. uuid: %s", uuid)

	var item LicenseServer.License

	query := fmt.Sprintf(
		`SELECT uuid, hardware_parameters, is_active, activated_on FROM license WHERE uuid = '%s'`, uuid)

	if err := r.db.Get(&item, query); err != nil {
		return item, nil
	}

	return item, nil
}
