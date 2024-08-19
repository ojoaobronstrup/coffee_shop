package controllers

import (
	"coffee_shop/internal/db"
	"coffee_shop/internal/jwt"
	"coffee_shop/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var informatedData models.LoginData

	err := c.BindJSON(&informatedData)
	if err != nil {
		log.Println("error on read the params data", err)
		c.JSON(500, "ERROR")
		return
	}

	db, err := db.DatabaseConnection()
	if err != nil {
		log.Println("error on establish connection with database", err)
		c.JSON(500, "ERROR")
		return
	}

	rows, err := db.Query("SELECT image_source, name FROM users WHERE username = ? AND password = ?", informatedData.Username, informatedData.Password)
	if err != nil {
		log.Println("error on get user from db", err)
		c.JSON(500, "ERROR")
		return
	}

	var imageSource string
	var name string

	if rows.Next() {
		err := rows.Scan(&imageSource, &name)
		if err != nil {
			log.Println("error on read the user credentials", err)
			c.JSON(500, "ERROR")
			return
		}
	}

	token, err := jwt.GenerateJWTToken(c)
	if err != nil {
		c.JSON(500, "ERROR")
		return
	}

	c.JSON(200, gin.H{
		"jwtToken":    token,
		"imageSource": imageSource,
		"name":        name,
	})
}
