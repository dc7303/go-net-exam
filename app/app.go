package app

import (
	"net/http"

	"go-net-exam/app/service"
)

type App struct{}

func (a *App) Run(addr string) error {
	initRoutes()
	return http.ListenAndServe(addr, nil)
}

func initRoutes() {
	http.HandleFunc("/", service.Index)
	http.HandleFunc("/getparam", service.GetParameter)
}
