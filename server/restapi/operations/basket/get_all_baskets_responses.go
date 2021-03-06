// Code generated by go-swagger; DO NOT EDIT.

package basket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/itimofeev/auchan/server/models"
)

// GetAllBasketsOKCode is the HTTP code returned for type GetAllBasketsOK
const GetAllBasketsOKCode int = 200

/*GetAllBasketsOK returns all available baskets

swagger:response getAllBasketsOK
*/
type GetAllBasketsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Basket `json:"body,omitempty"`
}

// NewGetAllBasketsOK creates GetAllBasketsOK with default headers values
func NewGetAllBasketsOK() *GetAllBasketsOK {

	return &GetAllBasketsOK{}
}

// WithPayload adds the payload to the get all baskets o k response
func (o *GetAllBasketsOK) WithPayload(payload []*models.Basket) *GetAllBasketsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all baskets o k response
func (o *GetAllBasketsOK) SetPayload(payload []*models.Basket) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllBasketsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Basket, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
