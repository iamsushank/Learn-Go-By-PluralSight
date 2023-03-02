package main

import (
	"github.com/iamsushank/Learn-Go-By-PluralSight/controller"
	"net/http"
)

func main() {
	/*u := models.User {
		ID: 2,
		FirstName: "Ram",
		LastName: "Kumar",
	}
	fmt.Println(u)*/

	controller.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
