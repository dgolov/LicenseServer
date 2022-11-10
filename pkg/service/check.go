package service

import (
	"fmt"
	"github.com/dgolov/LicenseServer/pkg/repository"
	"github.com/sirupsen/logrus"
	"time"
)

type CheckService struct {
	repo repository.License
}

func NewCheckLicense(repo repository.License) *CheckService {
	return &CheckService{repo: repo}
}

func (s *CheckService) CheckLicense(LicenseUuid string, HardwareParameters string) (int, error) {
	// Проверка лицензии
	logrus.Println("[Unit] CheckService - CheckLicense")

	if len(LicenseUuid) == 0 || len(HardwareParameters) == 0 {
		logrus.Error("Error input data. License uuid or hardware params is not found.")
		return 409, fmt.Errorf("error input data. License uuid or hardware params is not found")
	}

	licenseParam, err := s.repo.GetLicenseByUuid(LicenseUuid)
	if err != nil {
		logrus.Error("Error getting license by uuid. ", err.Error())
		return 409, fmt.Errorf("error getting license by uuid")
	}

	if licenseParam.IsActive != true {
		return 403, fmt.Errorf("license %s is not active", LicenseUuid)
	}

	if licenseParam.HardwareParameters != HardwareParameters {
		return 403, fmt.Errorf("license %s hardware parrametrs is not correct", LicenseUuid)
	}

	if licenseParam.ActivatedOn.Before(time.Now()) {
		return 403, fmt.Errorf("license %s expired", LicenseUuid)
	}

	return 200, nil
}
