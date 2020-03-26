package main

import (
	"fmt"
	"io"
	"net/http"
)

const userAgent = "cuelang.org/play/ playground snippet fetcher"

func handle(w http.ResponseWriter, r *http.Request) {
	cors(w)
	f := func(format string, args ...interface{}) {
		fmt.Fprintf(w, format, args...)
	}
	client := &http.Client{}
	if r.Method == "POST" {
		// Share
		url := fmt.Sprintf("https://play.golang.org/share")
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			f("Failed to create onwards GET URL: %v", err)
			return
		}
		req.Header.Add("User-Agent", userAgent)
		req.Body = r.Body
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			f("Failed in onward request: %v", err)
			return
		}
		io.Copy(w, resp.Body)
	} else {
		// Retrieve via the parameter id
		url := fmt.Sprintf("https://play.golang.org/p/%v.go", r.FormValue("id"))
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			f("Failed to create onwards GET URL: %v", err)
			return
		}
		req.Header.Add("User-Agent", userAgent)
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			f("Failed in onward request: %v", err)
			return
		}
		io.Copy(w, resp.Body)
	}
}

func main() {
	http.HandleFunc("/.netlify/functions/snippets", handle)
	serve()
}
