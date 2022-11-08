package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) check(c *gin.Context) {
	// Check license handler
	logrus.Println("Handler test")
}
