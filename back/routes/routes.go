package routes

import (
	"coffee_shop/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/login", controllers.Login)
}
