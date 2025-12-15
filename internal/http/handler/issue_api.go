package handler

import (
	"migrant-id/internal/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleIssue(c *gin.Context) {
	var req domain.MigrantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	tokenData, err := h.service.IssueID(req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	resp := domain.IssueResponse{
		Token:          tokenData.Token,
		ExpiresAt:      tokenData.ExpiresAt.Format(time.RFC822),
		QRCodeImageURL: "/api/qr/" + tokenData.PassportID,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) HandleQRImage(c *gin.Context) {
	passportID := c.Param("passportID")
	if passportID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing PassportID"})
		return
	}

	tokenData, err := h.service.GetTokenByPassport(passportID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found or expired"})
		return
	}

	c.Data(http.StatusOK, "image/png", tokenData.QRCodePNG)
}
