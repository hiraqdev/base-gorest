package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

const (
	// DefaultEnvPrefix is a constant variable used as prefix
	// for all used environment variables
	DefaultEnvPrefix = "GOREST"
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

// Env used to initialize Viper.  By default
// we are using viper to read configuration
// from environment variables
func Env() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(DefaultEnvPrefix)
}

// routes used to register all available routers
// parsing the data from Routers
func routes(h *mux.Router) http.Handler {
	for route, controller := range Routers {
		h.HandleFunc(route, controller.Handler).Methods(controller.Method)
	}

	return h
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
		jsonAPIHeaderFilterMiddleware,
	)

	srv := &http.Server{
		Handler:      engine,
		Addr:         options.Address,
		WriteTimeout: time.Duration(options.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(options.ReadTimeout) * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
