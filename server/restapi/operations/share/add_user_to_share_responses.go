// Code generated by go-swagger; DO NOT EDIT.

package share

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/itimofeev/auchan/server/models"
)

// AddUserToShareOKCode is the HTTP code returned for type AddUserToShareOK
const AddUserToShareOKCode int = 200

/*AddUserToShareOK returns created share

swagger:response addUserToShareOK
*/
type AddUserToShareOK struct {

	/*
	  In: Body
	*/
	Payload *models.Share `json:"body,omitempty"`
}

// NewAddUserToShareOK creates AddUserToShareOK with default headers values
func NewAddUserToShareOK() *AddUserToShareOK {

	return &AddUserToShareOK{}
}

// WithPayload adds the payload to the add user to share o k response
func (o *AddUserToShareOK) WithPayload(payload *models.Share) *AddUserToShareOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to share o k response
func (o *AddUserToShareOK) SetPayload(payload *models.Share) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToShareOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
