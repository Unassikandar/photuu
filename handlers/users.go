package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unassikandar/photuu/db"
)

func GetUsers(c *gin.Context) {
	// userId := c.Param("id")
	fmt.Printf("visited get user handler\n")

	for user, login := range db.Users {
		fmt.Printf("User: %s\n  hashedPass: %s\n  crsfTok: %s\n  sessionTok: %s\n",
			user, login.HashedPassword, login.CSRFToken, login.SessionToken)
	}
	c.JSON(200, "retrieved user")
}
