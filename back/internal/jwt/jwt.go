package jwt

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWTToken(c *gin.Context) (string, error) {
	err := godotenv.Load("C:/Users/joaog/DEV/coffee_shop/back/.env")
	if err != nil {
		log.Println("error on load the .env: ", err)
		return "", err
	}

	key := os.Getenv("KEY")
	var token *jwt.Token = jwt.New(jwt.SigningMethodHS256)

	stringToken, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("error on generate jwt token: ", err)
		return "", err
	}

	return stringToken, nil
}
