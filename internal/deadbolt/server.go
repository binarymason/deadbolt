package deadbolt

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type request struct {
	ip   string
	auth string
	path string
}

func (dblt *Deadbolt) defaultHandler(w http.ResponseWriter, r *http.Request) {
	rq := parseRequest(r)

	if rq.path != "/" {
		http.Error(w, "invalid route", http.StatusNotFound)
		return
	}
	w.Write([]byte("Deadbolt version: " + Version() + "\n"))
}

// sshdHandler updates a hosts sshd_config. Valid requests must meet this criteria:
//   * IP address is whitelisted in deadbolt.yml
//   * "Authorization: <secret>" must match deadbolt secret.
//
// The deadbolt secret is specified in deadbolt.yml or DEADBOLT_SECRET environment variable
//
// Handled routes:
// "/lock"   -> changes PermitRootLogin to "PermitRootLogin no"
// "/unlock" -> changes PermitRootLogin to "PermitRootLogin without-password"
func (dblt *Deadbolt) sshdHandler(w http.ResponseWriter, r *http.Request) {
	rq := parseRequest(r)
	if r.Method != http.MethodPost {
		http.Error(w, "invalid HTTP method. expected POST", http.StatusBadRequest)
		return
	}

	if !dblt.authorizedRequest(rq) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	permitRootLoginSetting := "no"

	if rq.path == "/unlock" {
		permitRootLoginSetting = "without-password"

	}

	if err := dblt.PermitRootLogin(permitRootLoginSetting); err != nil {
		errMessage := fmt.Sprintf("PermitRootLogin setting change failed: %v", err)
		log.Println("ERROR: ", errMessage)

		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(rq.path + "\n"))
}

func parseRequest(r *http.Request) *request {
	rq := request{
		ip:   strings.Split(r.RemoteAddr, ":")[0], // remove port
		auth: r.Header.Get("Authorization"),
		path: r.URL.Path,
	}

	return &rq
}

func (dblt *Deadbolt) authorizedRequest(r *request) bool {
	return validIP(r.ip, dblt.Whitelisted) && validAuth(r.auth, dblt.Secret)
}

func validIP(ip string, whitelisted []string) bool {
	for _, w := range whitelisted {
		if ip == w {
			return true
		}
	}

	return false
}

func validAuth(a, s string) bool {
	return a == s
}
