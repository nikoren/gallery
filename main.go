package main

import (
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("JOPA"))
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
