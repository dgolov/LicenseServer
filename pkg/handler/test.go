package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) test(c *gin.Context) {
	// HealthCheck
	log.Println("Handler test")
}
