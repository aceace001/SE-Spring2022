package users

import (
	"fmt"
	// "io/ioutil"
	"main/domain/users"
	"main/services"
	"main/utils/errors"

	// "main/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

const (
	SecretKey = "qwe123"
)

// unit test starts
func Add(a int, b int) int {
    return a + b
}

func Mul(a int, b int) int {
    return a * b
}

// ends 

//  users register an account
func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invald json body")
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// users login the account
func Login(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		err := errors.NewInternalServerErrror("login failed")
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, result)
}

// get users' information
func Get(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerErrror("could not retrive cookie")
		c.JSON(getErr.Status, getErr)
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		restErr := errors.NewInternalServerErrror("error parsing cookie")
		c.JSON(restErr.Status, restErr)
		return
	}
	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// log out the account
func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// check if email exist or not
func Email(c *gin.Context) {

}

// sprint3
// create homepage
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to ...",
	})
}

// type Book struct {
// 	ID     int64  `json:"id" gorm:"primary_key"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// 	// Content string `json:"content"`
// }

// var DB *gorm.DB

// func ConnectDatabase() {
// 	database, err := gorm.Open("Sqlites", "test.db")

// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	database.AutoMigrate(&Book{})

// 	DB = database
// }

type CreatePostsInput struct {
	Title   string `json:"title"  binding:"required"`
	Author  string `json:"author" binding:"required"`
	// Content string `json:"content" binding:"required"`
}

type UpdatePostsInput struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
}

// create posts
func PostHomePage(c *gin.Context) {
	fmt.Println("create posts")
	// Validate input
	var input CreatePostsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"message": book})
}

// search posts
func FindPosts(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"messges": books})
}

// search a post
func FindAPost(c *gin.Context) {
	var book models.Book 
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": book})
}
// update posts
func UpdatePosts(c *gin.Context) {
	var book models.Book 
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} 

	// validate input 
	var input UpdatePostsInput 
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"message": book})
}

// delete posts
func DeletePosts(c *gin.Context) {
	var book models.Book 
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return 
	}
	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
