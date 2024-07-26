package controllers

import (
	"coffee_shop/db"
	"coffee_shop/jwt"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db, err := db.DatabaseConnection()
	if err != nil {
		log.Println("error on establish connection with database: ", err)
		c.JSON(500, "ERROR")
		return
	}

	rows, err := db.Query("SELECT username, password FROM users")
	if err != nil {
		log.Println("error on get user from db: ", err)
		c.JSON(500, "ERROR")
		return
	}

	var username, password string

	for rows.Next() {

		err := rows.Scan(&username, &password)
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

	c.JSON(200, username)
}
