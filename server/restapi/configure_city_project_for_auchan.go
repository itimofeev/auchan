// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/itimofeev/auchan/server/restapi/operations"
	"github.com/itimofeev/auchan/server/restapi/operations/goods"
	"github.com/itimofeev/auchan/server/restapi/operations/share"
	"github.com/itimofeev/auchan/server/restapi/operations/user"
)

//go:generate swagger generate server --target ../server --name CityProjectForAuchan --spec ../tools/swagger.yml

func configureFlags(api *operations.CityProjectForAuchanAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CityProjectForAuchanAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-Auth-Token" header is set
	api.AuthTokenAuth = AuthFunc

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.GoodsAddGoodsToBasketHandler = goods.AddGoodsToBasketHandlerFunc(func(params goods.AddGoodsToBasketParams) middleware.Responder {
		return middleware.NotImplemented("operation goods.AddGoodsToBasket has not yet been implemented")
	})
	api.ShareAddUserToShareHandler = share.AddUserToShareHandlerFunc(func(params share.AddUserToShareParams) middleware.Responder {
		return middleware.NotImplemented("operation share.AddUserToShare has not yet been implemented")
	})
	api.BasketCreateBasketHandler = BasketCreateBasketHandler
	api.BasketGetAllBasketsHandler = BasketGetAllBasketsHandler
	api.GoodsGetAllGoodsInBasketHandler = goods.GetAllGoodsInBasketHandlerFunc(func(params goods.GetAllGoodsInBasketParams) middleware.Responder {
		return middleware.NotImplemented("operation goods.GetAllGoodsInBasket has not yet been implemented")
	})
	api.ShareGetAllSharesForBasketHandler = share.GetAllSharesForBasketHandlerFunc(func(params share.GetAllSharesForBasketParams) middleware.Responder {
		return middleware.NotImplemented("operation share.GetAllSharesForBasket has not yet been implemented")
	})
	api.ProductGetProductsByParamsHandler = ProductGetProductsByParamsHandler
	api.UserGetCurrentUserHandler = UserGetCurrentUserHandler
	api.UserLoginUserHandler = UserLoginUserHandler
	api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
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
