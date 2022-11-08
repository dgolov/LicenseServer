package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) test(c *gin.Context) {
	// HealthCheck
	logrus.Println("Handler test")
}
