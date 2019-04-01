package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type AuthTokenClaim struct {
	*jwt.StandardClaims
	User
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		var user User
		err := json.Unmarshal(r.Body, &user)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			fmt.Println("Error decoding json in registratoin", err)
			return
		}
		err = saveUser(user)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			fmt.Println("Error to save new user", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "method not supported", http.StatusNotAcceptable)
	}

}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		var user User
		err := json.Unmarshal(r.Body, &user)

		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			fmt.Println("Error decoding json in registratoin", err)
			return
		}

		model, err := getUser(user.Email)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			fmt.Println("not found user", err)
			return
		}

		err = model.checkPassword(user.Password)
		if err != nil {
			http.Error(w, "bad password", http.StatusBadRequest)
			fmt.Println("password is wrong", err, user)
			return
		}

	default:
		http.Error(w, "method not supported", http.StatusNotAcceptable)
	}
}
