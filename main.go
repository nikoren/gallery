package main

import (
	"headers"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set(headers.ContentTypeTextHtml())
	w.Write([]byte("<h1>JOPA</jopa>"))
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
