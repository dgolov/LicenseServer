package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
