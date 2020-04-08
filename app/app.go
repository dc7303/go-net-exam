package app

import (
	"net/http"

	"go-net-exam/app/service"
)

// run server
func Run(addr string) error {
	mux := initRoutes()
	return http.ListenAndServe(addr, mux)
}

// init mux
func initRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", service.Index)
	mux.HandleFunc("/bar", service.Bar)
	mux.Handle("/foo", &service.FooHandler{})

	return mux
}
