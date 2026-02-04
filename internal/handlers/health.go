package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthResponse struct {
	Status string `json:"status"`
}

// @Summary Health Check
// @Description Provide API health status
// @Success 200
// @Router /health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, healthResponse{Status: "ok"})
}
