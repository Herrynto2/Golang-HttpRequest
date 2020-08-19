package main

import (
	"github.com/Herrynto2/Golang-HttpRequest/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	routers.Setup()
}

func main() {

	server := gin.Default()
	server.GET("/index", Data)

	server.Run(":9001")
}
