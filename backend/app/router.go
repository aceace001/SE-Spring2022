package app

import (
	"main/controller/users"
	"time"

	"github.com/gin-contrib/cors"
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
}
