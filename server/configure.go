package server

import (
	"github.com/gorilla/mux"
	"github.com/roessland/withoutings/middleware"
	"github.com/roessland/withoutings/server/app"
	"github.com/roessland/withoutings/server/handlers"
	"net/http"
	"time"
)

func Configure(app *app.App) *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.Health(app))
	r.HandleFunc("/auth/callback", handlers.Callback(app))
	r.HandleFunc("/", handlers.HomePage(app))
	r.Use(
		middleware.Logging(app),
	)

	return &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9094",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
