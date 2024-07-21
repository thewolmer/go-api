package main

import (
	"github.com/thewolmer/go-api/config"
	"github.com/thewolmer/go-api/models"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}
func main() {

	config.DB.AutoMigrate(&models.Post{})

}
