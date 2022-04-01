package main

import (
	"fmt"
	"main/app"

	// "github.com/gin-gonic/gin"
)

// sprint3
// func HomePage(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Welcome to ...",
// 	})
// }

// func PostHomePage(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Post Home Page",
// 	})
// }
func main() {
	// sprint 3
	fmt.Println("*****hello*******")
	// r := gin.Default()
	// r.GET("/", HomePage)
	// r.POST("/", PostHomePage)
	// r.Run()

	// sprint 1
	app.StartApplication()
}
