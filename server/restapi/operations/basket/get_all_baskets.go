// Code generated by go-swagger; DO NOT EDIT.

package basket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAllBasketsHandlerFunc turns a function with the right signature into a get all baskets handler
type GetAllBasketsHandlerFunc func(GetAllBasketsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllBasketsHandlerFunc) Handle(params GetAllBasketsParams) middleware.Responder {
	return fn(params)
}

// GetAllBasketsHandler interface for that can handle valid get all baskets params
type GetAllBasketsHandler interface {
	Handle(GetAllBasketsParams) middleware.Responder
}

// NewGetAllBaskets creates a new http.Handler for the get all baskets operation
func NewGetAllBaskets(ctx *middleware.Context, handler GetAllBasketsHandler) *GetAllBaskets {
	return &GetAllBaskets{Context: ctx, Handler: handler}
}

/*GetAllBaskets swagger:route GET /basket basket getAllBaskets

get all users baskets

*/
type GetAllBaskets struct {
	Context *middleware.Context
	Handler GetAllBasketsHandler
}

func (o *GetAllBaskets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllBasketsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
