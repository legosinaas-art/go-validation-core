package handler

import (
	"fmt"
	"html/template"

	"migrant-id/internal/service"
)

var templates *template.Template

type Handler struct {
	service *service.MigrantService
}

func NewHandler(service *service.MigrantService) *Handler {
	// Инициализация шаблонов один раз при создании хендлера
	if templates == nil {
		// Шаблоны ожидаются в директории templates относительно корня проекта
		tmpl, err := template.ParseFiles(
			"templates/check.html",
			"templates/check_result.html",
			"templates/issue.html",
			"templates/issue_result.html",
		)
		if err != nil {
			panic(fmt.Sprintf("failed to parse templates: %v", err))
		}
		templates = tmpl
	}

	return &Handler{service: service}
}
