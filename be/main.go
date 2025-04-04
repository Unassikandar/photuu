package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/auth/google/callback", googleCallbackHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	verifier := oauth2.GenerateVerifier()
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	tok, err := googleOauthConfig.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		fmt.Println("Error exchanging code: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := googleOauthConfig.Client(ctx, tok)
	fmt.Println("Fetched client")
	client.Get("http://localhost:8080/done")
}
