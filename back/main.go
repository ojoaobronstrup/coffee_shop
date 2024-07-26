package main

import (
	"coffee_shop/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.Routes(r)
	r.Run(":8081")
}
