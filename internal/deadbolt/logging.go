package deadbolt

import (
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// logRequest is a closure that allows logging the response for Deadbolt's http server.
func logRequest(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := loggingResponseWrapper(w)
		wrappedHandler.ServeHTTP(lrw, req)
		statusCode := lrw.statusCode
		log.Println(req.RemoteAddr, req.Method, req.URL, statusCode, http.StatusText(statusCode))
	})
}

func loggingResponseWrapper(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}