package domain

import "time"

type MigrantRequest struct {
	PassportID string `json:"passport_id"`
	FullName   string `json:"full_name"`
}

type MigrantToken struct {
	PassportID string    `json:"passport_id"`
	Token      string    `json:"token"`
	ExpiresAt  time.Time `json:"expires_at"`
	QRCodePNG  []byte    `json:"-"`
}
type IssueResponse struct {
	Token          string `json:"token"`
	ShortToken     string `json:"short_token"`
	ExpiresAt      string `json:"expires_at"`
	QRCodeImageURL string `json:"qr_code_image_url"`
}
