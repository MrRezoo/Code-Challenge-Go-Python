package routers

import (
	"github.com/MrRezoo/code-challenge/api/handlers"
	"net/http"
)

func HealthRouter(mux *http.ServeMux) {
	handler := handlers.NewHealthHandler()

	mux.HandleFunc("/health/", handler.Health)
}
