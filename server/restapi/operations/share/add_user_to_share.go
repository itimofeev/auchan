// Code generated by go-swagger; DO NOT EDIT.

package share

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// AddUserToShareHandlerFunc turns a function with the right signature into a add user to share handler
type AddUserToShareHandlerFunc func(AddUserToShareParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddUserToShareHandlerFunc) Handle(params AddUserToShareParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddUserToShareHandler interface for that can handle valid add user to share params
type AddUserToShareHandler interface {
	Handle(AddUserToShareParams, interface{}) middleware.Responder
}

// NewAddUserToShare creates a new http.Handler for the add user to share operation
func NewAddUserToShare(ctx *middleware.Context, handler AddUserToShareHandler) *AddUserToShare {
	return &AddUserToShare{Context: ctx, Handler: handler}
}

/*AddUserToShare swagger:route POST /basket/{basketId}/share share addUserToShare

add user to share

*/
type AddUserToShare struct {
	Context *middleware.Context
	Handler AddUserToShareHandler
}

func (o *AddUserToShare) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddUserToShareParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddUserToShareBody add user to share body
// swagger:model AddUserToShareBody
type AddUserToShareBody struct {

	// email
	// Required: true
	Email *string `json:"email"`
}

// Validate validates this add user to share body
func (o *AddUserToShareBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddUserToShareBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("share"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddUserToShareBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddUserToShareBody) UnmarshalBinary(b []byte) error {
	var res AddUserToShareBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
