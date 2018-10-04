package app

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/hiraqdev/base-gorest/app/helper/jsonapi"
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

// jsonAPIHeaderFilterMiddleware used to check current request headers
// we should only serve request with valid jsonapi header request
// ref: http://jsonapi.org/format/#content-negotiation
func jsonAPIHeaderFilterMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentContentType := r.Header.Get("Content-Type")
		currentAcceptType := r.Header.Get("Accept")

		logHeader := log.WithFields(log.Fields{
			"Content-Type": currentContentType,
			"Accept":       currentAcceptType,
		})

		switch {
		case currentContentType != "application/vnd.api+json" || currentContentType == "":
			logHeader.Error("Invalid content type")
			errorContentType := jsonapi.NewError(http.StatusUnsupportedMediaType, "Unsupported Content Type", "Cannot continue your request")
			jsonAPIErrors := jsonapi.NewErrors(errorContentType)

			j, _ := jsonAPIErrors.GetErrors()

			w.Header().Set("Content-Type", "application/vnd.api+json")
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write(j)
		case currentAcceptType != "application/vnd.api+json" && currentAcceptType != "":
			logHeader.Error("Invalid accept type")
			errorAcceptType := jsonapi.NewError(http.StatusNotAcceptable, "Unsupported Access Type", "Cannot continue your request")
			jsonAPIErrors := jsonapi.NewErrors(errorAcceptType)

			j, _ := jsonAPIErrors.GetErrors()

			w.Header().Set("Content-Type", "application/vnd.api+json")
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(j)
		default:
			h.ServeHTTP(w, r)
		}
	})
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
