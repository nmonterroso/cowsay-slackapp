package main

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/nmonterroso/cowsay-slackapp/restapi/middleware"
	"github.com/nmonterroso/cowsay-slackapp/restapi/operations"
	"github.com/nmonterroso/cowsay-slackapp/restapi/responders"
	"math/rand"
	"time"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureAPI(api *operations.CowsaySlackappAPI) http.Handler {
	rand.Seed(time.Now().UTC().UnixNano())

	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	api.CowsayHandler = operations.CowsayHandlerFunc(responders.CowsayResponder)

	api.OauthRedirectHandler = operations.OauthRedirectHandlerFunc(responders.OauthRedirectResponder)

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return &middleware.FormContentTypeConversionMiddleware{handler}
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
