package main

import (
	"fmt"
	"time"

	"github.com/akshay-jadhav-cis/broadband-speed-tester/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello ")
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		fmt.Printf("Method=%s Origin=%q\n",
			c.Request.Method,
			c.GetHeader("Origin"),
		)
		c.Next()
	})
	config := cors.Config{
		AllowOrigins: []string{
			"https://speed-tester-wheat.vercel.app",
			"http://localhost:5000",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))
	router.Use(cors.New(config))
	router.SetTrustedProxies([]string{"127.0.0.1"})
	routes.AllRoutes(router)
	router.Run(":8080")
}
