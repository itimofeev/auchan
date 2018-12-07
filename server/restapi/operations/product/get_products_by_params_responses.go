// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/itimofeev/auchan/server/models"
)

// GetProductsByParamsOKCode is the HTTP code returned for type GetProductsByParamsOK
const GetProductsByParamsOKCode int = 200

/*GetProductsByParamsOK successful operation

swagger:response getProductsByParamsOK
*/
type GetProductsByParamsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Product `json:"body,omitempty"`
}

// NewGetProductsByParamsOK creates GetProductsByParamsOK with default headers values
func NewGetProductsByParamsOK() *GetProductsByParamsOK {

	return &GetProductsByParamsOK{}
}

// WithPayload adds the payload to the get products by params o k response
func (o *GetProductsByParamsOK) WithPayload(payload []*models.Product) *GetProductsByParamsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products by params o k response
func (o *GetProductsByParamsOK) SetPayload(payload []*models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsByParamsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Product, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetProductsByParamsNotFoundCode is the HTTP code returned for type GetProductsByParamsNotFound
const GetProductsByParamsNotFoundCode int = 404

/*GetProductsByParamsNotFound Not found

swagger:response getProductsByParamsNotFound
*/
type GetProductsByParamsNotFound struct {
}

// NewGetProductsByParamsNotFound creates GetProductsByParamsNotFound with default headers values
func NewGetProductsByParamsNotFound() *GetProductsByParamsNotFound {

	return &GetProductsByParamsNotFound{}
}

// WriteResponse to the client
func (o *GetProductsByParamsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}