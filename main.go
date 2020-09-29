package main

import (
	"net/http"

	"github.com/realpamisa/gmail-extractor/oauth"
	"github.com/realpamisa/gmail-extractor/pages"
)

func main() {
	http.HandleFunc("/", pages.HandleMain)
	http.HandleFunc("/login", oauth.HandleGoogleLogin)
	http.HandleFunc("/endpoint", oauth.HandleGoogleCallback)
	http.ListenAndServe(":8080", nil)
}
