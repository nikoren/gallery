package controllers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
)



func  debugRequest(w http.ResponseWriter, r *http.Request) {
	// print the request on DEBUG
	reqDump, err := httputil.DumpRequest(r, false)
	if err != nil {
		log.Errorf("Couldn't dump request content: %s", err.Error())
	}
	log.Debug(string(reqDump))
}
