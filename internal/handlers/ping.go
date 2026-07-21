package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PingCheck(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"status":    "Pong",
		"timeStamp": time.Now().UnixMilli(),
	})
}
