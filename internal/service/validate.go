package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MigrantStatus struct {
	IsValid   bool   `json:"is_valid"`
	Message   string `json:"message"`
	Name      string `json:"name,omitempty"`
	ExpiresAt string `json:"expires_at,omitempty"`
}

func (s *MigrantService) ValidateID(tokenString string) MigrantStatus {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	})

	if err != nil || !token.Valid {
		return MigrantStatus{IsValid: false, Message: "TOKEN INVALID/FAKE: Подпись недействительна или просрочен."}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return MigrantStatus{IsValid: false, Message: "TOKEN INVALID: Неверный формат данных."}
	}

	passportID, _ := claims["sub"].(string)
	fullName, _ := claims["name"].(string)

	var expTime time.Time
	if expTimestamp, ok := claims["exp"].(float64); ok {
		expTime = time.Unix(int64(expTimestamp), 0)
	} else if expTimestampInt, ok := claims["exp"].(int64); ok {
		expTime = time.Unix(expTimestampInt, 0)
	}

	// В. Проверка статуса (Имитация МВД - может быть аннулирован после выдачи)
	if !s.CheckGos(passportID) {
		return MigrantStatus{IsValid: false, Message: "STATUS ILLEGAL: Учет аннулирован/Депорт (Проверка МВД)."}
	}

	// Г. Финальная проверка: Всё ОК
	return MigrantStatus{
		IsValid:   true,
		Message:   "ДОКУМЕНТ ПОДЛИННЫЙ. Доступ разрешен.",
		Name:      fullName,
		ExpiresAt: expTime.Format(time.RFC822),
	}
}
