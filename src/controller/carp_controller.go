package controller

import (
	"net/http"
	"sagasugoi/model"

	"github.com/gin-gonic/gin"
)

func GetCarps(c *gin.Context) {
	result := model.ReadAllCarps()
	c.JSON(200, result)
}

func PostCarp(c *gin.Context) {
	var json model.Carp
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.CreateCarp(json)
	result := model.ReadAllCarps()
	c.JSON(200, result)
}
