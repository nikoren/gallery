package main

import (
	"gallery/views"
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
	sighnUpView *views.View
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
	handleIfErr("Couldn't render view", homeView.Render(w, nil))
}

func contacHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Contact requested")
	w.Header().Set(headers.ContentTypeTextHtml())
	handleIfErr("Couldn't render contact view", contactView.Render(w, nil))
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Signup requested")
	w.Header().Set(headers.ContentTypeTextHtml())
	handleIfErr("Couldn't render signup view", sighnUpView.Render(w, nil))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Warn("Unknown path was requested: ", r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
	handleIfErr("Couldn't render 404", notFoundView.Render(w, nil))
}

func main() {
	// Temp
	homeView = views.NewView(
		"bootstrap", LayoutsDir, "views/home.gohtml")
	contactView = views.NewView(
		"bootstrap", LayoutsDir,"views/contact.gohtml")
	notFoundView = views.NewView(
		"bootstrap", LayoutsDir,"views/404.gohtml")
	sighnUpView = views.NewView(
		"bootstrap", LayoutsDir,"views/signup.gohtml")

	// Routing
	r := mux.Router{}
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/home", homeHandler)
	r.HandleFunc("/contact", contacHandler)
	r.HandleFunc("/signup", signUpHandler)

	// Serving
	handleIfErr("Coudn't start server", http.ListenAndServe(":3000", &r))
}
