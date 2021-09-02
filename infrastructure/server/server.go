package server

import "net/http"

type Server interface {
	SetRoute(path string, handler AppHandler)
	Run() error
}

type AppHandler func(http.ResponseWriter, *http.Request) error
