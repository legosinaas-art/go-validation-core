package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleCheck(c *gin.Context) {
	if err := templates.ExecuteTemplate(c.Writer, "check.html", nil); err != nil {
		c.Status(http.StatusInternalServerError)
	}
}

func (h *Handler) HandleCheckResult(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		c.Redirect(http.StatusFound, "/")
		return
	}

	status := h.service.ValidateID(tokenString)

	if err := templates.ExecuteTemplate(c.Writer, "check_result.html", status); err != nil {
		c.Status(http.StatusInternalServerError)
	}
}
