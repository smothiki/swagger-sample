package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/deis/swagger-sample/pkg/swagger/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.SwaggerSampleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerSampleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.OauthAuth = func(token string, scopes []string) (interface{}, error) {
		fmt.Println(token)
		return "admin", nil
		//return nil, errors.NotImplemented("oauth2 bearer auth (oauth) has not yet been implemented")
	}

	api.GetKingHandler = operations.GetKingHandlerFunc(func(params operations.GetKingParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .GetKing has not yet been implemented")
	})
	api.GetPingHandler = operations.GetPingHandlerFunc(func(params operations.GetPingParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .GetPing has not yet been implemented")
	})
	api.GetClusterByIDHandler = operations.GetClusterByIDHandlerFunc(func(params operations.GetClusterByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetClusterByID has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
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
