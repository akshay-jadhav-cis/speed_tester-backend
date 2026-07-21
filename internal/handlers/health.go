package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthStatus(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"server":  "BroadBand Speed Tester",
		"backend": "Go",
		"version": "1.0.0.0",
	})
}
