package service

import (
	"github.com/dgolov/LicenseServer/pkg/repository"
)

type License interface {
	CheckLicense(LicenseUuid string, HardwareParameters string) (int, error)
}

type Service struct {
	License
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		License: NewCheckLicense(repos.License),
	}
}
