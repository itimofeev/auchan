// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/itimofeev/auchan/server/restapi/operations/basket"
	"github.com/itimofeev/auchan/server/restapi/operations/goods"
	"github.com/itimofeev/auchan/server/restapi/operations/product"
	"github.com/itimofeev/auchan/server/restapi/operations/share"
	"github.com/itimofeev/auchan/server/restapi/operations/user"
)

// NewCityProjectForAuchanAPI creates a new CityProjectForAuchan instance
func NewCityProjectForAuchanAPI(spec *loads.Document) *CityProjectForAuchanAPI {
	return &CityProjectForAuchanAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		GoodsAddGoodsToBasketHandler: goods.AddGoodsToBasketHandlerFunc(func(params goods.AddGoodsToBasketParams) middleware.Responder {
			return middleware.NotImplemented("operation GoodsAddGoodsToBasket has not yet been implemented")
		}),
		ShareAddUserToShareHandler: share.AddUserToShareHandlerFunc(func(params share.AddUserToShareParams) middleware.Responder {
			return middleware.NotImplemented("operation ShareAddUserToShare has not yet been implemented")
		}),
		BasketCreateBasketHandler: basket.CreateBasketHandlerFunc(func(params basket.CreateBasketParams) middleware.Responder {
			return middleware.NotImplemented("operation BasketCreateBasket has not yet been implemented")
		}),
		BasketGetAllBasketsHandler: basket.GetAllBasketsHandlerFunc(func(params basket.GetAllBasketsParams) middleware.Responder {
			return middleware.NotImplemented("operation BasketGetAllBaskets has not yet been implemented")
		}),
		GoodsGetAllGoodsInBasketHandler: goods.GetAllGoodsInBasketHandlerFunc(func(params goods.GetAllGoodsInBasketParams) middleware.Responder {
			return middleware.NotImplemented("operation GoodsGetAllGoodsInBasket has not yet been implemented")
		}),
		ShareGetAllSharesForBasketHandler: share.GetAllSharesForBasketHandlerFunc(func(params share.GetAllSharesForBasketParams) middleware.Responder {
			return middleware.NotImplemented("operation ShareGetAllSharesForBasket has not yet been implemented")
		}),
		ProductGetProductsByParamsHandler: product.GetProductsByParamsHandlerFunc(func(params product.GetProductsByParamsParams) middleware.Responder {
			return middleware.NotImplemented("operation ProductGetProductsByParams has not yet been implemented")
		}),
		UserGetUserByNameHandler: user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation UserGetUserByName has not yet been implemented")
		}),
		UserLoginUserHandler: user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserLoginUser has not yet been implemented")
		}),
		UserLogoutUserHandler: user.LogoutUserHandlerFunc(func(params user.LogoutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserLogoutUser has not yet been implemented")
		}),

		// Applies when the "X-Auth-Token" header is set
		AuthTokenAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (AuthToken) X-Auth-Token from header param [X-Auth-Token] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*CityProjectForAuchanAPI This is a server for collaborative shopping.
 */
type CityProjectForAuchanAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// AuthTokenAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Auth-Token provided in the header
	AuthTokenAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// GoodsAddGoodsToBasketHandler sets the operation handler for the add goods to basket operation
	GoodsAddGoodsToBasketHandler goods.AddGoodsToBasketHandler
	// ShareAddUserToShareHandler sets the operation handler for the add user to share operation
	ShareAddUserToShareHandler share.AddUserToShareHandler
	// BasketCreateBasketHandler sets the operation handler for the create basket operation
	BasketCreateBasketHandler basket.CreateBasketHandler
	// BasketGetAllBasketsHandler sets the operation handler for the get all baskets operation
	BasketGetAllBasketsHandler basket.GetAllBasketsHandler
	// GoodsGetAllGoodsInBasketHandler sets the operation handler for the get all goods in basket operation
	GoodsGetAllGoodsInBasketHandler goods.GetAllGoodsInBasketHandler
	// ShareGetAllSharesForBasketHandler sets the operation handler for the get all shares for basket operation
	ShareGetAllSharesForBasketHandler share.GetAllSharesForBasketHandler
	// ProductGetProductsByParamsHandler sets the operation handler for the get products by params operation
	ProductGetProductsByParamsHandler product.GetProductsByParamsHandler
	// UserGetUserByNameHandler sets the operation handler for the get user by name operation
	UserGetUserByNameHandler user.GetUserByNameHandler
	// UserLoginUserHandler sets the operation handler for the login user operation
	UserLoginUserHandler user.LoginUserHandler
	// UserLogoutUserHandler sets the operation handler for the logout user operation
	UserLogoutUserHandler user.LogoutUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *CityProjectForAuchanAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CityProjectForAuchanAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CityProjectForAuchanAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CityProjectForAuchanAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CityProjectForAuchanAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CityProjectForAuchanAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CityProjectForAuchanAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CityProjectForAuchanAPI
func (o *CityProjectForAuchanAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AuthTokenAuth == nil {
		unregistered = append(unregistered, "XAuthTokenAuth")
	}

	if o.GoodsAddGoodsToBasketHandler == nil {
		unregistered = append(unregistered, "goods.AddGoodsToBasketHandler")
	}

	if o.ShareAddUserToShareHandler == nil {
		unregistered = append(unregistered, "share.AddUserToShareHandler")
	}

	if o.BasketCreateBasketHandler == nil {
		unregistered = append(unregistered, "basket.CreateBasketHandler")
	}

	if o.BasketGetAllBasketsHandler == nil {
		unregistered = append(unregistered, "basket.GetAllBasketsHandler")
	}

	if o.GoodsGetAllGoodsInBasketHandler == nil {
		unregistered = append(unregistered, "goods.GetAllGoodsInBasketHandler")
	}

	if o.ShareGetAllSharesForBasketHandler == nil {
		unregistered = append(unregistered, "share.GetAllSharesForBasketHandler")
	}

	if o.ProductGetProductsByParamsHandler == nil {
		unregistered = append(unregistered, "product.GetProductsByParamsHandler")
	}

	if o.UserGetUserByNameHandler == nil {
		unregistered = append(unregistered, "user.GetUserByNameHandler")
	}

	if o.UserLoginUserHandler == nil {
		unregistered = append(unregistered, "user.LoginUserHandler")
	}

	if o.UserLogoutUserHandler == nil {
		unregistered = append(unregistered, "user.LogoutUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CityProjectForAuchanAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CityProjectForAuchanAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "AuthToken":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.AuthTokenAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *CityProjectForAuchanAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *CityProjectForAuchanAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *CityProjectForAuchanAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CityProjectForAuchanAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the city project for auchan API
func (o *CityProjectForAuchanAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CityProjectForAuchanAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/basket/{basketId}/goods"] = goods.NewAddGoodsToBasket(o.context, o.GoodsAddGoodsToBasketHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/basket/{basketId}/share"] = share.NewAddUserToShare(o.context, o.ShareAddUserToShareHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/basket"] = basket.NewCreateBasket(o.context, o.BasketCreateBasketHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/basket"] = basket.NewGetAllBaskets(o.context, o.BasketGetAllBasketsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/basket/{basketId}/goods"] = goods.NewGetAllGoodsInBasket(o.context, o.GoodsGetAllGoodsInBasketHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/basket/{basketId}/share"] = share.NewGetAllSharesForBasket(o.context, o.ShareGetAllSharesForBasketHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/product"] = product.NewGetProductsByParams(o.context, o.ProductGetProductsByParamsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/{email}"] = user.NewGetUserByName(o.context, o.UserGetUserByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/login"] = user.NewLoginUser(o.context, o.UserLoginUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/logout"] = user.NewLogoutUser(o.context, o.UserLogoutUserHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CityProjectForAuchanAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CityProjectForAuchanAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CityProjectForAuchanAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CityProjectForAuchanAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
