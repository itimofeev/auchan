// Code generated by go-swagger; DO NOT EDIT.

package goods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/itimofeev/auchan/server/models"
)

// AddGoodsToBasketOKCode is the HTTP code returned for type AddGoodsToBasketOK
const AddGoodsToBasketOKCode int = 200

/*AddGoodsToBasketOK returns current state of goods

swagger:response addGoodsToBasketOK
*/
type AddGoodsToBasketOK struct {

	/*
	  In: Body
	*/
	Payload *models.Goods `json:"body,omitempty"`
}

// NewAddGoodsToBasketOK creates AddGoodsToBasketOK with default headers values
func NewAddGoodsToBasketOK() *AddGoodsToBasketOK {

	return &AddGoodsToBasketOK{}
}

// WithPayload adds the payload to the add goods to basket o k response
func (o *AddGoodsToBasketOK) WithPayload(payload *models.Goods) *AddGoodsToBasketOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add goods to basket o k response
func (o *AddGoodsToBasketOK) SetPayload(payload *models.Goods) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGoodsToBasketOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGoodsToBasketNotFoundCode is the HTTP code returned for type AddGoodsToBasketNotFound
const AddGoodsToBasketNotFoundCode int = 404

/*AddGoodsToBasketNotFound product or basket not found

swagger:response addGoodsToBasketNotFound
*/
type AddGoodsToBasketNotFound struct {
}

// NewAddGoodsToBasketNotFound creates AddGoodsToBasketNotFound with default headers values
func NewAddGoodsToBasketNotFound() *AddGoodsToBasketNotFound {

	return &AddGoodsToBasketNotFound{}
}

// WriteResponse to the client
func (o *AddGoodsToBasketNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
