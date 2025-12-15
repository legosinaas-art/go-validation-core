package router

import (
	"migrant-id/internal/http/handler"

	"github.com/gin-gonic/gin"
)

// NewRouter configures all HTTP routes and middleware for the service.
func NewRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	web := r.Group("/")
	registerWebRoutes(web, h)

	api := r.Group("/api")
	registerAPIRoutes(api, h)

	return r
}

func registerWebRoutes(rg *gin.RouterGroup, h *handler.Handler) {
	rg.GET("/", h.HandleCheck)
	rg.GET("/check/result", h.HandleCheckResult)

	rg.GET("/issue", h.HandleIssueForm)
	rg.POST("/issue/process", h.HandleIssueProcess)
}

func registerAPIRoutes(rg *gin.RouterGroup, h *handler.Handler) {
	rg.GET("/validate", h.HandleValidate)
	rg.POST("/issue", h.HandleIssue)
	rg.GET("/qr/:passportID", h.HandleQRImage)
}
