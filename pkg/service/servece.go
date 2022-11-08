package service

import "github.com/dgolov/LicenseServer/pkg/repository"

type License interface {
}

type Service struct {
	License
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
