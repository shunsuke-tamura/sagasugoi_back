package controller

import (
	"net/http"
	"sagasugoi/model"

	"github.com/gin-gonic/gin"
)

func GetSASToken(c *gin.Context) {
	res, err := model.GeterateSASToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
