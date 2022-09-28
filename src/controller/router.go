package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/types")
	{
		api.GET("/", Index)
		api.GET("/:id", ShowOneType)
		api.POST("/create", CreateType)
		api.POST("/edit", EditType)
		api.POST("/delete", DeleteType)
	}
	return r
}
