package restapi

import (
	"net/http"

	errors "github.com/go-swagger/go-swagger/errors"
	httpkit "github.com/go-swagger/go-swagger/httpkit"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"

	"restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.ConfigServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ConfigServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	api.DeleteDataSomeKeyPathHandler = operations.DeleteDataSomeKeyPathHandlerFunc(func(params operations.DeleteDataSomeKeyPathParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteDataSomeKeyPath has not yet been implemented")
	})
	api.GetDataSomeKeyPathHandler = operations.GetDataSomeKeyPathHandlerFunc(func(params operations.GetDataSomeKeyPathParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetDataSomeKeyPath has not yet been implemented")
	})
	api.PostDataSomeKeyPathHandler = operations.PostDataSomeKeyPathHandlerFunc(func(params operations.PostDataSomeKeyPathParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostDataSomeKeyPath has not yet been implemented")
	})
	api.PutDataSomeKeyPathHandler = operations.PutDataSomeKeyPathHandlerFunc(func(params operations.PutDataSomeKeyPathParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutDataSomeKeyPath has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
