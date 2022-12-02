package main

import (
	"sagasugoi/controller"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		panic("Error loading env file")
	}
	r := controller.GetRouter()
	r.Run()
}

// package main

// import (
// 	"sagasugoi/controller"
// )

// func main() {

// 	r := controller.GetRouter()
// 	r.Run(":8080")
// router := gin.Default()
// router.GET("/", func(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello は",
// 	})
// })
// db.Init()
// router.Run()
// db.Close()()
// }
