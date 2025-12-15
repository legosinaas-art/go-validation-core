package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleValidate(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'token' query parameter"})
		return
	}

	status := h.service.ValidateID(tokenString)

	code := http.StatusOK
	if !status.IsValid {
		code = http.StatusForbidden
	}

	c.JSON(code, status)
}


