package routers

import "github.com/gin-gonic/gin"

//Data golang di visual studio
func Data(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello!",
	})
}
