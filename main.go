package main

import (
	"views/models"

	"github.com/gin-gonic/gin"
)

func posting(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.GET("/index", models.GetData)
	router.POST("/upload", posting)
	router.Run(":9001")
}
