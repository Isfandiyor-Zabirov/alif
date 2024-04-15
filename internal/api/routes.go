package api

import (
	"alif/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// RunApp launches all routes
// @title Alif bank tech task
// @version 1.0
// @description Open API документация для проекта
// @host localhost
// @BasePath /
func RunApp() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	v1 := router.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to project"})
	})

	notificationRoutes := v1.Group("/notifications")
	notificationRoutes.POST("", handlers.CreateNotification)
	notificationRoutes.GET("", handlers.GetAllNotifications)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "reason": "Page not found"})
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("Cannot run the app:", err.Error())
	}
}

// CORSMiddleware solve cors problem by adding headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Refresh-Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}
