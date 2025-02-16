package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // Marking this as production release to disable debugging mode

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"}) // Only trust these proxy servers

	router.GET("/healthz", healthzHandler)
	router.GET("/api/version", apiVersionHandler)

	router.Run("localhost:7002")
}

func healthzHandler(c *gin.Context) {
	c.String(http.StatusOK, "I am alive")
}

func apiVersionHandler(c *gin.Context) {
	data := gin.H{
		"version":  "v1.0.0",
		"app_name": "app_002",
	}
	c.IndentedJSON(http.StatusOK, data)
}
