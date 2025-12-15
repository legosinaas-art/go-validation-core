package repository

import (
	"errors"
	"migrant-id/internal/domain"
)

func (r *InMemoryRepo) GetTokenByPassport(id string) (*domain.MigrantToken, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.storage[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &val, nil
}
