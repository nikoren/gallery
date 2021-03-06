package controllers

import (
	"fmt"
	"gallery/views"
	"net/http"
	"shortcuts"
	log "github.com/sirupsen/logrus"
)

// UsersC is a controller that can be served as main entry point for all Users
// related business logic such as creating,updating, deleting and any any other
// types of modifications of the users
type UsersC struct {
	View *views.View
}

// This function is used for initial setup of UsersC controller
// you should only use it once when you create the controller,
// it will fail if UsersC.View template is not parsed correctly
//
func NewUsersC() *UsersC {
	return &UsersC{
		View: views.NewView("bootstrap","views/users/create.gohtml"),
	}
}

type decodedForm struct {
	Email    string `schema:"form-email"`
	Password string `schema:"form-password"`
}

// Create  is used to handle all the business logic (implementation details)
// related to creation of new users objects
func (u *UsersC) Create(w http.ResponseWriter, r *http.Request) {

	// GET
	if r.Method == http.MethodGet {
		w.Header().Set(shortcuts.ContentTypeTextHtml())
		u.View.Render(w, nil)
	}

	// POST
	if r.Method == http.MethodPost {
		var df decodedForm
		if err := shortcuts.ParseForm(r, &df); err != nil{
			log.Panicf("Couldn't parse form: %s", err.Error())
		}
		// User the extracted form values
		_,_ = fmt.Fprintln(w, "Submited form:", df)
		log.Infof("New user signed up %s", df.Email)
	}
}
