package routes

import (
	"net/http"

	"github.com/akshay-jadhav-cis/broadband-speed-tester/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AllRoutes(router *gin.Engine) {
	router.GET("/", func(g *gin.Context) {
		g.JSON(http.StatusOK, gin.H{
			"message": "Backend SuccessFully Running",
		})
	})
	router.GET("/health", handlers.HealthStatus)
	router.GET("/ping", handlers.PingCheck)
	router.POST("/upload", handlers.UploadCheck)
}
