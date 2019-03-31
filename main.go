package main

import (
	"gallery/controllers"
	"gallery/views"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"shortcuts"
)

const (
	Port       = 3000
	LayoutsDir = "views/layouts"
)

var (
	homeView     *views.View
	contactView  *views.View
	notFoundView *views.View
)

func handleIfErr(msg string, err error) {
	if err != nil {
		log.Fatalf("MSG: %s, ERROR: %s", msg, err)
	}
}

func init() {
	// setup logging
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
}

func main() {
	// Register controllers - business logic goes here
	usersC := controllers.NewUsersC()
	staticC := controllers.NewStaticC()

	// Routing - map views/controllers to http routes
	r := mux.Router{}
	r.NotFoundHandler = staticC.NotFoundV
	r.Handle("/home", staticC.HomeV).Methods("GET")
	r.Handle("/contact", staticC.ContactV).Methods("Get")
	r.HandleFunc("/users/create", usersC.Create).Methods("GET","POST")

	// middlewares
	r.Use(shortcuts.InspectRequestsMiddleware)

	// Start the server...
	log.Info("Starting serving on :3000")
	handleIfErr("Coudn't start server", http.ListenAndServe(":3000", &r))
}
