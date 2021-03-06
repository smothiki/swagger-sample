package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*GetKingOK server ping success

swagger:response getKingOK
*/
type GetKingOK struct {
}

// NewGetKingOK creates GetKingOK with default headers values
func NewGetKingOK() *GetKingOK {
	return &GetKingOK{}
}

// WriteResponse to the client
func (o *GetKingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*GetKingDefault unexpected error

swagger:response getKingDefault
*/
type GetKingDefault struct {
	_statusCode int
}

// NewGetKingDefault creates GetKingDefault with default headers values
func NewGetKingDefault(code int) *GetKingDefault {
	if code <= 0 {
		code = 500
	}

	return &GetKingDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get king default response
func (o *GetKingDefault) WithStatusCode(code int) *GetKingDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get king default response
func (o *GetKingDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *GetKingDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
}
