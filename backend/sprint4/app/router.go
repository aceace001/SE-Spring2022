package app

import (
	"main/controller/users"
	"time"

	"github.com/gin-contrib/cors"

	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func mapUrls() {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("/api/register", users.Register)
	router.POST("/api/login", users.Login)
	router.GET("/api/user", users.Get)
	router.GET("/api/logout", users.Logout)
	router.GET("/api/email", users.Email)

	// chat
	// router.GET("/", users.ChatRoom)
	m := melody.New()
	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "home.html")
	})

	router.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	// sprint3
	// router.GET("/", users.HomePage)
	router.POST("/post", users.PostHomePage)

	router.GET("/post", users.FindPosts)
	router.GET("/post/:id", users.FindAPost)
	router.PATCH("/post/:id", users.UpdatePosts)
	router.DELETE("/post/:id", users.DeletePosts)

}
