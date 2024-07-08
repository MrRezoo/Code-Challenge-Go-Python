package middlewares

import (
	"log"
	"net/http"
)

const (
	colorReset       = "\033[0m"
	colorCyan        = "\033[36m"
	colorRed         = "\033[31m"
	bgColorLightBlue = "\033[44m"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(customWriter, r)

		statusCode := customWriter.statusCode
		color := colorCyan
		if statusCode >= 400 && statusCode < 500 {
			color = colorRed
		}

		log.Printf("%s%s %s %d %s%s%s", color, r.Method, r.URL.Path, statusCode, bgColorLightBlue, colorReset, colorReset)
	})
}
