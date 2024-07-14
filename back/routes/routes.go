package routes

import (
	"coffee_shop/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/login", controllers.Login)
}
