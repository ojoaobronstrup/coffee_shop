package controllers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWTToken(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load("./.env")
	if err != nil {
		slog.Error("error on load the .env")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("error on configure the token"))
	}

	key := os.Getenv("KEY")
	var token *jwt.Token = jwt.New(jwt.SigningMethodES256)

	stringToken, err := token.SignedString(key)
	if err != nil {
		slog.Error("error on load the .env")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("error on configure the token"))
	}

	cookie := http.Cookie{
		Name:  "jwt_token",
		Value: stringToken,
	}

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("success on configure the token"))
}

func Login(w http.ResponseWriter, r *http.Request) {

}
