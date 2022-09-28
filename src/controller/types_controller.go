package controller

import (
	"fmt"
	"sagasugoi/model"
	"strconv"

	// "strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	datas := model.GetAll()
	fmt.Println(datas)
	c.JSON(200, datas)

}
func ShowOneType(c *gin.Context) {
	id := c.Params.ByName("id")
	data := model.GetOne(id)
	c.JSON(200, data)
}
func CreateType(c *gin.Context) {
	id := c.PostForm("id")
	size, _ := strconv.ParseFloat(c.PostForm("size"), 64)
	r, _ := strconv.Atoi(c.PostForm("r"))
	g, _ := strconv.Atoi(c.PostForm("g"))
	b, _ := strconv.Atoi(c.PostForm("b"))

	data := model.Types{Id: id, Size: size, R: r, G: g, B: b}
	// data := model.Types{c.JSON}
	data.Create()
	c.JSON(200, data)
}

func EditType(c *gin.Context) {
	id := c.PostForm("id")
	data := model.GetOne(id)
	size, _ := strconv.ParseFloat(c.PostForm("size"), 64)
	data.Size = size
	r, _ := strconv.Atoi(c.PostForm("r"))
	data.R = r
	g, _ := strconv.Atoi(c.PostForm("g"))
	data.G = g
	b, _ := strconv.Atoi(c.PostForm("b"))
	data.B = b
	data.Update()
	c.JSON(200, data)
}

func DeleteType(c *gin.Context) {
	id := c.PostForm("id")
	data := model.GetOne(id)
	data.Delete()
	c.JSON(200, "success")
}
