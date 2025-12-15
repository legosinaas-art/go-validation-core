package repository

import (
	"migrant-id/internal/domain"
	"sync"
)

type InMemoryRepo struct {
	storage map[string]domain.MigrantToken
	mu      sync.RWMutex
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		storage: make(map[string]domain.MigrantToken),
	}
}

type Repository interface {
	SaveToken(migrant domain.MigrantToken) error
	GetTokenByPassport(passportID string) (*domain.MigrantToken, error)
}
