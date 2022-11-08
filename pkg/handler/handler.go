package handler

import (
	"github.com/dgolov/LicenseServer/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// All routes
	router := gin.New()

	main := router.Group("/")
	{
		main.POST("/test", h.test)
		main.POST("/check", h.check)
	}

	return router
}
