package main

import (
	"coffee_shop/routes"
	"net/http"
)

func main() {
	routes.Routes()
	http.ListenAndServe("localhost:8081", nil)
}
