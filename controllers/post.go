package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thewolmer/go-api/config"
	"github.com/thewolmer/go-api/models"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	fmt.Println(body)
	post := models.Post{Title: body.Title, Body: body.Body}
	result := config.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}
