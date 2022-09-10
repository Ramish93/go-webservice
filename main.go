package main

import (
	"net/http"

	"github.com/ramish93/go-webservice/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}