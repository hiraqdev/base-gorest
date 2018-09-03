package app

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

// Middleware used to intercept request and manage
// internal state api engine
type Middleware func(http.Handler) http.Handler

// chainMiddleware used to register all available middlewares
// to the main api engine
func chainMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}

	return h
}

// jsonAPIHeaderResponseMiddleware used to manipulate response header
// type, using default jsonapi.org standard: application/vnd.api+json
// ref: http://jsonapi.org/format/#content-negotiation-servers
func jsonAPIHeaderResponseMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.api+json")
		h.ServeHTTP(w, r)
	})
}

// optionLogginMiddleware used to show available server options to stdout
func optionLogginMiddleware(options ServerConfig) func(h http.Handler) http.Handler {
	log.WithFields(log.Fields{
		"Server Address": options.Address,
		"Read Timeout":   options.ReadTimeout,
		"Write Timeout":  options.WriteTimeout,
	}).Info("Server configurations")

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})
	}
}

// logginMiddleware used to create default logging message
// for every http access.
// ref: http://www.gorillatoolkit.org/pkg/handlers#LoggingHandler
func loggingMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

// recoveryMiddleware used to recovers system from panic
// ref: http://www.gorillatoolkit.org/pkg/handlers#RecoveryHandler
func recoveryMiddleware(h http.Handler) http.Handler {
	return handlers.RecoveryHandler()(h)
}
