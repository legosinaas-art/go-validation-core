package service

import "migrant-id/internal/domain"

func (s *MigrantService) GetTokenByPassport(passportID string) (*domain.MigrantToken, error) {
	return s.repo.GetTokenByPassport(passportID)
}
