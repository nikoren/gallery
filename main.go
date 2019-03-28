package main

import (
	"gallery/views"
	"gallery/controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"headers"
	"net/http"
)

const (
	Port = 3000
	LayoutsDir = "views/layouts"
)

var (
	homeView    *views.View
	contactView *views.View
	notFoundView *views.View
)

func handleIfErr(msg string, err error ){
	if err != nil{
		log.Fatalf("MSG: %s, ERROR: %s", msg, err)
	}
}

func init(){
	// setup logging
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Home requested")
	w.Header().Set(headers.ContentTypeTextHtml())
	homeView.Render(w, nil)
}

func contacHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Contact requested")
	w.Header().Set(headers.ContentTypeTextHtml())
	contactView.Render(w, nil)
}


func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Warn("Unknown path was requested: ", r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
	notFoundView.Render(w, nil)
}

func main() {

	// Register views  - just for rendering things , not for business logic
	homeView = views.NewView(
		"bootstrap", LayoutsDir, "views/home.gohtml")
	contactView = views.NewView(
		"bootstrap", LayoutsDir,"views/contact.gohtml")
	notFoundView = views.NewView(
		"bootstrap", LayoutsDir,"views/404.gohtml")

	// Register controllers - business logic goes here
	usersC := controllers.NewUsersC()

	// Routing - map views/controllers to http routes
	r := mux.Router{}
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/home", homeHandler)
	r.HandleFunc("/contact", contacHandler)
	r.HandleFunc("/users/create", usersC.Create)

	// Start the server...
	log.Info("Starting serving on :3000")
	handleIfErr("Coudn't start server", http.ListenAndServe(":3000", &r))
}
