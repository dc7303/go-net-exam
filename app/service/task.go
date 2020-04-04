package service

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func GetParameter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Query())
}
