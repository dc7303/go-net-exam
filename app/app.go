package app

import (
	"net/http"

	"go-net-exam/app/service"

	"github.com/gorilla/mux"
)

// run server
func Run(addr string) error {
	mux := initRoutes()
	return http.ListenAndServe(addr, mux)
}

// init mux
func initRoutes() *mux.Router {
	mux := mux.NewRouter()
	// mux := http.NewServeMux()
	mux.HandleFunc("/", service.Index)
	mux.HandleFunc("/bar", service.Bar)
	mux.Handle("/foo", &service.FooHandler{})
	mux.Handle(
		"/fileserver/",
		http.StripPrefix(
			"/fileserver/",
			http.FileServer(http.Dir("public")),
		),
	)
	mux.HandleFunc("/uploads", service.UploadsHanlder)
	mux.HandleFunc("/users/{id}", service.UsersHandler).Methods("GET", "POST")

	return mux
}
