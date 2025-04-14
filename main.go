package main

import (
	"encoding/base64"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetLevel(log.ErrorLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func handleEncode(input string, c *gin.Context) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	c.JSON(http.StatusOK, gin.H{"encoded": encoded, "status": http.StatusOK})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(gin.Recovery()) // "recover from any panics", write 500 if any

	r.NoRoute(func(c *gin.Context) {
		// log this event as it could be an attempt to break in...
		log.Errorln("requested URL path not found:", c.Request.URL.Path)

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "requested resource not found", "status": http.StatusNotFound})
	})

	r.GET("/spenc/:string", func(c *gin.Context) {
		input := c.Param("string") + "\n" // if reached, a string is provided, otherwise gin would have returned a 404 before

		handleEncode(input, c)
	})

	r.GET("/encode/:string", func(c *gin.Context) {
		input := c.Param("string") // if reached, a string is provided, otherwise gin would have returned a 404 before

		handleEncode(input, c)
	})

	r.GET("/decode/:string", func(c *gin.Context) {
		input := c.Param("string") // gin returns a 404 if no string is provided

		decoded, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid base64 string", "status": http.StatusBadRequest})
			return
		}

		c.JSON(http.StatusOK, gin.H{"decoded": string(decoded), "status": http.StatusOK})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "alive and kicking", "status": http.StatusOK})
	})

	r.OPTIONS("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"routes": []string{
				"GET /spenc/:string - Base64-encodes the string with a newline appended",
				"GET /encode/:string - Base64-encodes the string",
				"GET /decode/:string - Base64-decodes the string",
				"GET /health - Health check endpoint",
			},
			"status": http.StatusOK,
		})
	})

	// set port via PORT environment variable
	r.Run() // default port is 8080
}
