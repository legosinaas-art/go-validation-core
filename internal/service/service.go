package service

import "migrant-id/internal/repository"

type MigrantService struct {
	repo   repository.Repository
	secret []byte
}

func NewMigrantService(repo repository.Repository, secret []byte) *MigrantService {
	return &MigrantService{
		repo:   repo,
		secret: secret,
	}
}
