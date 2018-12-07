// Code generated by go-swagger; DO NOT EDIT.

package goods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/itimofeev/auchan/server/models"
)

// GetAllGoodsInBasketOKCode is the HTTP code returned for type GetAllGoodsInBasketOK
const GetAllGoodsInBasketOKCode int = 200

/*GetAllGoodsInBasketOK returns goods in basket

swagger:response getAllGoodsInBasketOK
*/
type GetAllGoodsInBasketOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Goods `json:"body,omitempty"`
}

// NewGetAllGoodsInBasketOK creates GetAllGoodsInBasketOK with default headers values
func NewGetAllGoodsInBasketOK() *GetAllGoodsInBasketOK {

	return &GetAllGoodsInBasketOK{}
}

// WithPayload adds the payload to the get all goods in basket o k response
func (o *GetAllGoodsInBasketOK) WithPayload(payload []*models.Goods) *GetAllGoodsInBasketOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all goods in basket o k response
func (o *GetAllGoodsInBasketOK) SetPayload(payload []*models.Goods) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllGoodsInBasketOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Goods, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
