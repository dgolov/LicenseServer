package LicenseServer

import "time"

type License struct {
	// License model
	Id                 int       `json:"-"`
	Uuid               string    `json:"uuid"`
	HardwareParameters string    `json:"hardware_parameters"`
	IsActive           bool      `json:"is_active"`
	ActivatedOn        time.Time `json:"activated_on"`
	ExpiratedOn        time.Time `json:"expirated_on"`
}
