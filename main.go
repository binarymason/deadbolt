package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const DEADBOLT_VERSION = "201907081400"

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/unlock", deadboltHandler)
	http.HandleFunc("/lock", deadboltHandler)

	port := ":8080"
	fmt.Println("listening on port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path != "/" {
		http.Error(w, "invalid route", http.StatusNotFound)
		return
	}
	w.Write([]byte("Deadbolt version: " + DEADBOLT_VERSION + "\n"))
}

func deadboltHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if r.Method != http.MethodPost {
		http.Error(w, "invalid HTTP method. expected POST", http.StatusBadRequest)
		return
	}
	w.Write([]byte(path + "\n"))
}

func lockSSHConfig(cfg string) string {
	return generateSSHConfig("lock", cfg)
}

func unlockSSHConfig(cfg string) string {
	return generateSSHConfig("unlock", cfg)
}

func generateSSHConfig(m, cfg string) string {
	lines := strings.Split(cfg, "\n")
	var result []string
	for _, line := range lines {
		result = append(result, toggle(m, line))
	}

	return strings.Join(result, "\n")
}

func toggle(m, s string) string {
	setting := "PermitRootLogin"
	match, _ := regexp.Match(`PermitRootLogin`, []byte(s))

	if !match {
		return s
	}

	switch m {
	case "lock":
		return setting + " no"
	case "unlock":
		return setting + " yes"
	default:
		panic("unhandled toggle method: " + m)

	}
}
