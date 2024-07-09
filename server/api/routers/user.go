package routers

import (
	"github.com/MrRezoo/code-challenge/api/handlers"
	. "github.com/MrRezoo/code-challenge/api/helpers"
	"net/http"
)

func UserRouter(mux *http.ServeMux) {
	handler := handlers.NewUserHandler()
	mux.HandleFunc("/user/all/", MethodHandler(http.MethodGet, handler.Users))
	mux.HandleFunc("/user/register/", MethodHandler(http.MethodPost, handler.RegisterUser))
	mux.HandleFunc("/user/deposit/", MethodHandler(http.MethodPatch, handler.Deposit))
	mux.HandleFunc("/user/withdraw/", MethodHandler(http.MethodPatch, handler.Withdraw))
	mux.HandleFunc("/user/bet/", MethodHandler(http.MethodPost, handler.Bet))
}
