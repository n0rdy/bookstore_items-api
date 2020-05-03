package application

import (
	"github.com/gorilla/mux"
	"github.com/n0rdy/bookstore_users-api/logger"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Addr: ":8081",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	logger.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
