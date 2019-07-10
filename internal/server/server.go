package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/binarymason/deadbolt/internal/config"
	"github.com/binarymason/deadbolt/internal/routes"
)

const DEADBOLT_VERSION = "201907081400"

type Server struct {
	router routes.Router
}

func New(p *string) *Server {
	s := Server{}

	s.router = routes.Router{
		Version: DEADBOLT_VERSION,
		Config:  config.Load(*p),
	}

	return &s
}

func (s *Server) Serve() {
	router := &s.router

	http.HandleFunc("/", router.Default)
	http.HandleFunc("/unlock", router.Deadbolt)
	http.HandleFunc("/lock", router.Deadbolt)

	port := router.Port()

	fmt.Println("listening on port", port)

	if err := http.ListenAndServe(port, logRequest(http.DefaultServeMux)); err != nil {
		panic(err)
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func logRequest(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)
		statusCode := lrw.statusCode
		log.Println(req.RemoteAddr, req.Method, req.URL, statusCode, http.StatusText(statusCode))
	})
}
