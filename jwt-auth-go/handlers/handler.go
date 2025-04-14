package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/auth"
)

var hardCodedEmail = "email@123.com"
var hardCodedPassword = "12345"

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var u UserInfo
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if u.Email != hardCodedEmail || u.Password != hardCodedPassword {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}
	tokenString, err := auth.CreateToken(u.Email)
	if err != nil {
		http.Error(w, "Token not generated", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)

}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("Missing authorization header")
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	err := auth.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	fmt.Fprint(w, "Welcome! You are authorized")
}
