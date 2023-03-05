package controller

import (
	"github.com/iamsushank/Learn-Go-By-PluralSight/models"
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

/*
implementing Handler interface

	type handler interface {
		serveHTTP(ResponseWriter, *Request)
	}
*/
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	users := models.GetUser()
	encodeResponseAsJSON(users, w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	user, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(user, w)
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
