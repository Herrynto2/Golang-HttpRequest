package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// server.GET("/index", "")

	server.Run(":9001")
}
