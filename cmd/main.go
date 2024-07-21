package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thewolmer/go-api/config"
	"github.com/thewolmer/go-api/controllers"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.Run()
}
