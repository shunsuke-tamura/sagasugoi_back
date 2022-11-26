package controller

import (
	"sagasugoi/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
	}))

	r.GET("/healthcheck", Healthcheck)

	carpApi := r.Group("/carps")
	{
		carpApi.GET("/", GetCarps)
		carpApi.POST("/", PostCarp)
	}
	r.GET("/token_for_upload_image", GetSASToken)

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

func Healthcheck(c *gin.Context) {
	status := model.Health{
		Status: "OK",
	}
	c.JSON(200, status)
}
