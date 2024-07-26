package main

import (
	"coffee_shop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routes.Routes(r)
	r.Run(":8081")
}
