package handler

import (
	"github.com/dgolov/LicenseServer"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) check(c *gin.Context) {
	// Check license handler
	logrus.Println("Handler check")

	var input LicenseServer.License

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	logrus.Printf("[Input] Uuid: %s\n Hardware parameters: %s", input.Uuid, input.HardwareParameters)

	statusCode, err := h.service.CheckLicense(input.Uuid, input.HardwareParameters)
	if err != nil {
		NewErrorResponse(c, statusCode, err.Error())
	}

	logrus.Println(statusCode)
}
