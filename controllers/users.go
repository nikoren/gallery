package controllers

import (
	"gallery/views"
	log "github.com/sirupsen/logrus"
	"headers"
	"net/http"
)

// Users is a controller that can be served as main entry point for all Users
// related business logic such as creating,updating, deleting and any any other
// types of modifications of the users
type Users struct {
	View *views.View
}

// This function is used for initial setup of Users controller
// you should only use it once when you create the controller,
// it will fail if Users.View template is not parsed correctly
//
func NewUsersC() *Users {
	return &Users{
		View: views.NewView("bootstrap",
			"views/layouts", ""+
				"views/users/create.gohtml"),
	}
}

// Create  is used to handle all the business logic (implementation details)
// related to creation of new users objects
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	if log.GetLevel() == log.DebugLevel{
		debugRequest(w, r)
	}

	// GET
	if r.Method == http.MethodGet {
		w.Header().Set(headers.ContentTypeTextHtml())
		u.View.Render(w, nil)
	}

	// POST
	if r.Method == http.MethodPost {
		log.Warnf("Values submitted : email=%s, pass=%s",
			r.PostFormValue("form-email"),
			r.PostFormValue("form-password"),
		)
	}
}

