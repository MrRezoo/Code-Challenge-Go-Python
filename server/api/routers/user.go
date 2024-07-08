package routers

import (
	"github.com/MrRezoo/code-challenge/api/handlers"
	"net/http"
)

func UserRouter(mux *http.ServeMux) {
	handler := handlers.NewUserHandler()

	mux.HandleFunc("/user/all", handler.Users)
	mux.HandleFunc("/user/register", handler.RegisterUser)
	mux.HandleFunc("/user/deposit", handler.Deposit)
	mux.HandleFunc("/user/withdraw", handler.Withdraw)
	mux.HandleFunc("/user/bet", handler.Bet)
}
