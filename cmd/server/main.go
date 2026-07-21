package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/akshay-jadhav-cis/broadband-speed-tester/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello ")
	router := gin.Default()
	config := cors.Config{
		AllowOriginFunc: func(origin string) bool {
        return origin == "http://localhost:5000" ||
               strings.HasSuffix(origin, ".vercel.app")
    },
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))
	router.SetTrustedProxies([]string{"127.0.0.1"})
	routes.AllRoutes(router)
	router.Run(":8080")
}
