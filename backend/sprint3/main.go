package main

import (
	// "fmt"
	"main/app"

	"github.com/rahmanfadhil/gin-bookstore/models"

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
	// sprint 1
	models.ConnectDatabase()
	app.StartApplication()
}
