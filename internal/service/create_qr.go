package service

import (
	"errors"
	"migrant-id/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skip2/go-qrcode"
)

func (s *MigrantService) IssueID(req domain.MigrantRequest) (*domain.MigrantToken, error) {
	if !s.CheckGos(req.PassportID) {
		return nil, errors.New("верификация не пройдена")
	}

	existing, _ := s.repo.GetTokenByPassport(req.PassportID)
	if existing != nil {
		return nil, errors.New("ID уже создано пользователем")
	}

	// Token is valid for 30 days.
	expirationTime := time.Now().Add(time.Hour * 24 * 30)
	claims := jwt.MapClaims{
		"sub":  req.PassportID,
		"name": req.FullName,
		"exp":  expirationTime.Unix(),
		"iss":  "migrant_id_system",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secret)
	if err != nil {
		return nil, err
	}

	png, err := qrcode.Encode(tokenString, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	result := domain.MigrantToken{
		PassportID: req.PassportID,
		Token:      tokenString,
		ExpiresAt:  expirationTime,
		QRCodePNG:  png,
	}

	err = s.repo.SaveToken(result)
	return &result, err
}
