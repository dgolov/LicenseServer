package LicenseServer

import "time"

type License struct {
	// License model
	Id                 int       `json:"-" db:"id"`
	Uuid               string    `json:"license_uuid" db:"uuid" binding:"required"`
	HardwareParameters string    `json:"hardware_parameters" db:"hardware_parameters" binding:"required"`
	IsActive           bool      `json:"is_active" db:"is_active"`
	ActivatedOn        time.Time `json:"activated_on" db:"activated_on"`
	ExpiratedOn        time.Time `json:"expirated_on" db:"expirated_on"`
}
