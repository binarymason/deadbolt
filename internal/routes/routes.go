package routes

import (
	"net/http"
	"strings"

	"github.com/binarymason/deadbolt/internal/config"
	"github.com/binarymason/deadbolt/internal/validate"
)

type Router struct {
	Config  config.Config
	Version string
}

func (rtr *Router) Port() string {
	p := "8080"
	if rtr.Config.Port != "" {
		p = rtr.Config.Port
	}
	return ":" + p
}

func (rtr *Router) Default(w http.ResponseWriter, r *http.Request) {
	rq := parseRequest(r)

	if rq.path != "/" {
		http.Error(w, "invalid route", http.StatusNotFound)
		return
	}
	w.Write([]byte("Deadbolt version: " + rtr.Version + "\n"))
}

func (rtr *Router) Deadbolt(w http.ResponseWriter, r *http.Request) {
	rq := parseRequest(r)
	if r.Method != http.MethodPost {
		http.Error(w, "invalid HTTP method. expected POST", http.StatusBadRequest)
		return
	}

	if !validate.ValidRequest(rq.ip, rq.auth, rtr.Config) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(rq.path + "\n"))
}

type request struct {
	ip   string
	auth string
	path string
}

func parseRequest(r *http.Request) *request {
	rq := request{
		ip:   strings.Split(r.RemoteAddr, ":")[0], // remove port
		auth: r.Header.Get("Authorization"),
		path: r.URL.Path,
	}

	return &rq
}
