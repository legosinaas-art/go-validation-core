package repository

import (
	"errors"
	"migrant-id/internal/domain"
)

func (r *InMemoryRepo) SaveToken(m domain.MigrantToken) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storage[m.PassportID]; exists {
		return errors.New("аккаунт уже создан")
	}

	r.storage[m.PassportID] = m
	return nil
}
