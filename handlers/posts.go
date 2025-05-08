package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unassikandar/photuu/utils"
)

func CreatePostHandler(c *gin.Context) {
	fmt.Println("create user handler")
	c.JSON(200, "created user")
}

func GetPosts(c *gin.Context) {
	if err := utils.Authorize(c.Request); err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	username := c.Request.FormValue("username")
	fmt.Printf("CSRF validation successful! Welcome, %s", username)
}
