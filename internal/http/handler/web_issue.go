package handler

import (
	"migrant-id/framework"
	"migrant-id/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleIssueForm(c *gin.Context) {
	if err := templates.ExecuteTemplate(c.Writer, "issue.html", nil); err != nil {
		c.Status(http.StatusInternalServerError)
	}
}

func (h *Handler) HandleIssueProcess(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing form"})
		return
	}

	req := domain.MigrantRequest{
		PassportID: c.PostForm("passport_id"),
		FullName:   c.PostForm("full_name"),
	}

	tokenData, err := h.service.IssueID(req)
	if err != nil {
		data := struct {
			Error string
		}{
			Error: err.Error(),
		}
		if tmplErr := templates.ExecuteTemplate(c.Writer, "issue.html", data); tmplErr != nil {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	resp := domain.IssueResponse{
		Token:          tokenData.Token,
		ShortToken:     framework.ShortenToken(tokenData.Token),
		ExpiresAt:      tokenData.ExpiresAt.Format("02.01.2006 15:04"),
		QRCodeImageURL: "/api/qr/" + tokenData.PassportID,
	}

	if err := templates.ExecuteTemplate(c.Writer, "issue_result.html", resp); err != nil {
		c.Status(http.StatusInternalServerError)
	}
}


