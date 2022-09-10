package main

import (
	"fmt"

	"github.com/ramish93/go-webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Ramish",
		LastName: "Hassan",
	}
	fmt.Println(u)
}