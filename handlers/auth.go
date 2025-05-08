package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unassikandar/photuu/db"
	"github.com/unassikandar/photuu/utils"
)

func Register(c *gin.Context) {
	fmt.Println("visited login handler")

	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	fmt.Printf("username: %s\npassword: %s", username, password)

	if len(username) < 3 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, "Invalid username/password")
		return
	}

	if _, ok := db.Users[username]; ok {
		c.JSON(http.StatusConflict, "User already exists")
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	db.Users[username] = db.Login{
		HashedPassword: hashedPassword,
	}

	c.JSON(http.StatusCreated, "Registered user")
}

func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	user, ok := db.Users[username]
	if !ok || !utils.CheckPassword(password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, "Invalid username or password")
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	//TODO: check params
	c.SetCookie("session_token", sessionToken, 86400, "", "localhost", false, true)
	c.SetCookie("csrf_token", csrfToken, 86400, "", "localhost", false, false)

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	db.Users[username] = user

	c.JSON(http.StatusOK, "Logged in")
}

func Logout(c *gin.Context) {
	if err := utils.Authorize(c.Request); err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.SetCookie("session_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("csrf_token", "", -1, "/", "localhost", false, true)

	username := c.Request.FormValue("username")
	user, _ := db.Users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	db.Users[username] = user

	c.JSON(http.StatusOK, "Logged out!")
}
