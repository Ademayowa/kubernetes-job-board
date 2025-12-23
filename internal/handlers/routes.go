package handlers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Allow frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Metrics middleware
	server.Use(metricsMiddleware())

	server.GET("/jobs", getJobs)
	server.POST("/jobs", createJob)
	server.GET("/healthz", healthCheck)

	// Metrics endpoint
	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
