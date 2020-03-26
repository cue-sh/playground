// +build !netlify

package main

import (
	"log"
	"net/http"
)

func serve() {
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
