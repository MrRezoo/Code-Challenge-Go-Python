package api

import (
	"github.com/MrRezoo/code-challenge/api/middlewares"
	"github.com/MrRezoo/code-challenge/api/routers"
	"github.com/MrRezoo/code-challenge/config"
	"github.com/MrRezoo/code-challenge/db"
	"log"
	"net/http"
	"os"
)

func initServer() *http.ServeMux {
	cfg := config.GetConfig()

	db.InitRedis(&cfg.Redis)

	mux := http.NewServeMux()

	muxWithMiddlewares := http.NewServeMux()
	muxWithMiddlewares.Handle("/", middlewares.Logging(middlewares.CORS(mux)))
	{

		routers.HealthRouter(mux)
		routers.UserRouter(mux)
	}
	return muxWithMiddlewares
}

func ServeMux() {
	cfg := config.GetConfig()
	serveMux := initServer()
	log.Printf("⚠️ Server is started on localhost:%s ⚠️", cfg.Server.Port)
	err := http.ListenAndServe(":"+cfg.Server.Port, serveMux)
	if err != nil {
		log.Printf("Failed to start server on %s: %v", cfg.Server.Port, err)
		os.Exit(1)
	}
}
