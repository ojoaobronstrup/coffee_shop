package routes

import (
	"coffee_shop/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
}
