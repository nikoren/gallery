package controllers

import (
	"gallery/views"
)
//  StaticC is a static controller that is used as a catch all controller for things that do not fit
// into any other dedicated controller
type StaticC struct {
	HomeV *views.View
	ContactV *views.View
	NotFoundV *views.View

}

// NewStaticC creates new static controller
func NewStaticC() *StaticC{
	return &StaticC{
		HomeV:     views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactV:  views.NewView("bootstrap",  "views/static/contact.gohtml"),
		NotFoundV: views.NewView("bootstrap",  "views/static/notfound.gohtml"),
	}
}


