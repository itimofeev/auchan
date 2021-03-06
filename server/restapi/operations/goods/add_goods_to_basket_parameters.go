// Code generated by go-swagger; DO NOT EDIT.

package goods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddGoodsToBasketParams creates a new AddGoodsToBasketParams object
// no default values defined in spec.
func NewAddGoodsToBasketParams() AddGoodsToBasketParams {

	return AddGoodsToBasketParams{}
}

// AddGoodsToBasketParams contains all the bound params for the add goods to basket operation
// typically these are obtained from a http.Request
//
// swagger:parameters addGoodsToBasket
type AddGoodsToBasketParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	BasketID int64
	/*
	  In: body
	*/
	Goods AddGoodsToBasketBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddGoodsToBasketParams() beforehand.
func (o *AddGoodsToBasketParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBasketID, rhkBasketID, _ := route.Params.GetOK("basketId")
	if err := o.bindBasketID(rBasketID, rhkBasketID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body AddGoodsToBasketBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("goods", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Goods = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBasketID binds and validates parameter BasketID from path.
func (o *AddGoodsToBasketParams) bindBasketID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("basketId", "path", "int64", raw)
	}
	o.BasketID = value

	return nil
}
