package controllers

import (
	"coffee_shop/db"
	"coffee_shop/jwt"
	"coffee_shop/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var informatedData models.LoginData

	err := c.BindJSON(&informatedData)
	if err != nil {
		log.Println("error on read the body request: ", err)
		return
	}

	db, err := db.DatabaseConnection()
	if err != nil {
		log.Println("error on establish connection with database: ", err)
		c.JSON(500, "ERROR")
		return
	}

	rows, err := db.Query("SELECT username FROM users WHERE username = ? AND password = ?", informatedData.Username, informatedData.Password)
	if err != nil {
		log.Println("error on get user from db: ", err)
		c.JSON(500, "ERROR")
		return
	}

	var username string

	for rows.Next() {

		err := rows.Scan(&username)
		if err != nil {
			log.Println("error on read the user credentials: ", err)
			c.JSON(500, "ERROR")
			return
		}
	}

	err = jwt.GenerateJWTToken(c)
	if err != nil {
		c.JSON(500, "ERROR")
		return
	}

	c.JSON(200, gin.H{username: username})
}
