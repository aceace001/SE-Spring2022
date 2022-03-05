package main

import (
	"log"
	"main/controller/users"
	"net/http"
	"os"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string
	Password string
}

var (
	// router
	router http.Handler

	// customed request headers for token authorization and so on
	myHeaders = make(map[string]string, 0)

	logging *log.Logger
	// a default user variable for the next requests parameter
	user = User{
		UserName: "Valiben",
		Password: "123456",
	}

	// a default token for authorization
	myToken = "ssoiuoiu"

	tokenName = "x-xq5-jwt"
)

// set the router
func SetRouter(r http.Handler) {
	router = r
}

// set the log
func SetLog(l *log.Logger) {
	logging = l
}

// add custom request header
func AddHeader(key, value string) {
	myHeaders[key] = value
}

// receive the ordinary response
type OrdinaryResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

// a middleware function for easy authorization
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(tokenName)
		if token != myToken {
			// unauthorized, just return the unauthorized response
			c.JSON(http.StatusUnauthorized, gin.H{"errno": "-1", "errmsg": "unauthorized"})
			return
		} else {
			// authorized, come into the next function
			c.Next()
		}
	}
}

func init() {
	// initialize the router
	router := gin.Default()

	// a handler for getting static resources
	router.Static("/file", "./")

	// some handlers for post/put/delete requests
	router.POST("/api/login", users.Login)
	router.POST("/api/register", users.Register)
	router.GET("/api/email", users.Email)
	router.GET("/api/user", users.Get)

	//router.GET("/username/to/password", GetPasswordHandler)
	//router.GET("/get/age", GetAgeHandler)
	//router.PUT("/add/user", AddUserHandler)
	//router.DELETE("/delete/user", DeleteUserHandler)
	//router.POST("/upload", SaveFileHandler)

	// use a middleware function
	router.Use(Authorize())

	// set the router
	SetRouter(router)

	// set customized request headers
	AddHeader(tokenName, myToken)

	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	SetLog(newLog)
}

//get
func TestGetAgeHandler(t *testing.T) {
	// make request params
	param := make(map[string]interface{})
	param["user_name"] = user.UserName
	param["password"] = user.Password

	resp := OrdinaryResponse{}

	err := unitTest.TestHandlerUnMarshalResp("GET", "/api/user", "form", param, &resp)
	if err != nil {
		t.Errorf("TestGetAgeHandler: %v\n", err)
		return
	}

	if resp.Errno != "0" {
		t.Errorf("TestGetAgeHandler: response is not expected\n")
		return
	}
}

//get users.Email
func TestGetEmailHandler(t *testing.T) {
	// make request params
	param := make(map[string]interface{})
	param["user_name"] = user.UserName

	resp := OrdinaryResponse{}

	err := unitTest.TestHandlerUnMarshalResp("GET", "/api/email", "form", param, &resp)
	if err != nil {
		t.Errorf("TestGetPasswordHandler: %v\n", err)
		return
	}

	if resp.Errno != "0" {
		t.Errorf("TestGetPasswordHandler: response is not expected\n")
		return
	}
}

//users.Login
func TestLoginHandler(t *testing.T) {
	// make request params
	param := make(map[string]interface{})
	param["user_name"] = user.UserName
	param["password"] = user.Password

	resp := OrdinaryResponse{}

	err := unitTest.TestHandlerUnMarshalResp("POST", "/api/login", "form", param, &resp)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}

	if resp.Errno != "0" {
		t.Errorf("TestLoginHandler: response is not expected\n")
		return
	}
}

//put (register)
func TestAddUserHandler(t *testing.T) {
	resp := OrdinaryResponse{}

	err := unitTest.TestHandlerUnMarshalResp("POST", "/api/register", "form", user, &resp)
	if err != nil {
		t.Errorf("TestAddUserHandler: %v\n", err)
		return
	}
	if resp.Errno != "0" {
		t.Errorf("TestAddUserHandler: response is not expected\n")
		return
	}
}
