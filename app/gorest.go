package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ServerConfig used to setup options to
// run our api engine
// Fields :
// Address (string)
// ReadTimeout (int)
// WriteTimeout (int)
type ServerConfig struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
}

// Gorest used to initiate main application server
func Gorest(options ServerConfig) {
	router := routes(mux.NewRouter())
	engine := chainMiddleware(
		router,
		optionLogginMiddleware(options),
		loggingMiddleware,
		recoveryMiddleware,
		jsonAPIHeaderResponseMiddleware,
	)

	srv := &http.Server{
		Handler:      engine,
		Addr:         options.Address,
		WriteTimeout: time.Duration(options.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(options.ReadTimeout) * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
