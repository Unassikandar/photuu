package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unassikandar/photuu/handlers"
)

func New() {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.GET("/users", handlers.GetUsers)
	r.POST("/login", handlers.Login)
	r.POST("/logout", handlers.Logout)

	r.GET("/posts", handlers.GetPosts)

	fmt.Println("listening on port 8080......")
	r.Run(":8080")
}
