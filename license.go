package LicenseServer

type License struct {
	Id                 int    `json:"-"`
	Uuid               string `json:"uuid"`
	HardwareParameters string `json:"hardware_parameters"`
	IsActive           bool   `json:"is_active"`
	ActivatedOn        string `json:"activated_on"`
	ExpiratedOn        string `json:"expirated_on"`
}
