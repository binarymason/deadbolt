package routes

import (
	"net/http"

	"github.com/binarymason/go-deadbolt/internal/config"
)

type Router struct {
	Config  config.Config
	Version string
}

func (rtr *Router) Port() string {
	p := ":"
	if rtr.Config.Port != "" {
		return p + rtr.Config.Port
	}
	return p + "8080"
}

func (rtr *Router) Default(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path != "/" {
		http.Error(w, "invalid route", http.StatusNotFound)
		return
	}
	w.Write([]byte("Deadbolt version: " + rtr.Version + "\n"))
}

func (rtr *Router) Deadbolt(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if r.Method != http.MethodPost {
		http.Error(w, "invalid HTTP method. expected POST", http.StatusBadRequest)
		return
	}
	w.Write([]byte(path + "\n"))
}
