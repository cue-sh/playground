// +build netlify

package main

import (
	"log"
	"net/http"

	"github.com/apex/gateway"
)

func serve() {
	log.Fatal(gateway.ListenAndServe(":8080", nil))
}

func cors(w http.ResponseWriter) {
}
