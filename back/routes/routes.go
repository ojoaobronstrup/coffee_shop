package routes

import (
	"coffee_shop/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/jwt", controllers.GenerateJWTToken)
	http.HandleFunc("/login", controllers.Login)
}
