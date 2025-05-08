package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/unassikandar/photuu/db"
)

var authError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := db.Users[username]
	if !ok {
		println("no user")
		return authError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		fmt.Printf("st: %s\n", st.Value)
		fmt.Printf("user.SessionToken: %s\n", user.SessionToken)
		return authError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		println("no csrf")
		fmt.Printf("csrf: %s\n", csrf)
		fmt.Printf("user.csrf: %s\n", user.CSRFToken)
		return authError
	}

	return nil
}
