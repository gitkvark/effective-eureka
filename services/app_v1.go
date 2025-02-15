package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthz", healthzHandler)
	router.GET("/api/version", apiVersionHandler)

	router.Run("localhost:7001")
}

func healthzHandler(c *gin.Context) {
	c.String(http.StatusOK, "I am alive")
}

func apiVersionHandler(c *gin.Context) {
	data := gin.H{
		"version":  "v0.0.1",
		"app_name": "app_001",
	}
	c.IndentedJSON(http.StatusOK, data)
}
