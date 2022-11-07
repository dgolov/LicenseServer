package service

type Check struct {
}

func (c *Check) CheckLicense(LicenseUuid string, HardwareParameters string) (string, int) {
	// Проверка лицензии
	if len(LicenseUuid) == 0 || len(HardwareParameters) == 0 {
		return "Error input data. License uuid or hardware params is not found.", 409
	}

	return "Ok", 200
}
