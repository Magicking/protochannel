// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Magicking/protochannel/models"
)

// SignOffCommitOKCode is the HTTP code returned for type SignOffCommitOK
const SignOffCommitOKCode int = 200

/*SignOffCommitOK sign off commit o k

swagger:response signOffCommitOK
*/
type SignOffCommitOK struct {
}

// NewSignOffCommitOK creates SignOffCommitOK with default headers values
func NewSignOffCommitOK() *SignOffCommitOK {
	return &SignOffCommitOK{}
}

// WriteResponse to the client
func (o *SignOffCommitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*SignOffCommitDefault Unexpected error

swagger:response signOffCommitDefault
*/
type SignOffCommitDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSignOffCommitDefault creates SignOffCommitDefault with default headers values
func NewSignOffCommitDefault(code int) *SignOffCommitDefault {
	if code <= 0 {
		code = 500
	}

	return &SignOffCommitDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the sign off commit default response
func (o *SignOffCommitDefault) WithStatusCode(code int) *SignOffCommitDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the sign off commit default response
func (o *SignOffCommitDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the sign off commit default response
func (o *SignOffCommitDefault) WithPayload(payload *models.Error) *SignOffCommitDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sign off commit default response
func (o *SignOffCommitDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignOffCommitDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}