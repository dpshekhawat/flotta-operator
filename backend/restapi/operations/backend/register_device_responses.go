// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/project-flotta/flotta-operator/backend/models"
)

// RegisterDeviceOKCode is the HTTP code returned for type RegisterDeviceOK
const RegisterDeviceOKCode int = 200

/*RegisterDeviceOK Updated

swagger:response registerDeviceOK
*/
type RegisterDeviceOK struct {
}

// NewRegisterDeviceOK creates RegisterDeviceOK with default headers values
func NewRegisterDeviceOK() *RegisterDeviceOK {

	return &RegisterDeviceOK{}
}

// WriteResponse to the client
func (o *RegisterDeviceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RegisterDeviceUnauthorizedCode is the HTTP code returned for type RegisterDeviceUnauthorized
const RegisterDeviceUnauthorizedCode int = 401

/*RegisterDeviceUnauthorized Unauthorized

swagger:response registerDeviceUnauthorized
*/
type RegisterDeviceUnauthorized struct {
}

// NewRegisterDeviceUnauthorized creates RegisterDeviceUnauthorized with default headers values
func NewRegisterDeviceUnauthorized() *RegisterDeviceUnauthorized {

	return &RegisterDeviceUnauthorized{}
}

// WriteResponse to the client
func (o *RegisterDeviceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RegisterDeviceForbiddenCode is the HTTP code returned for type RegisterDeviceForbidden
const RegisterDeviceForbiddenCode int = 403

/*RegisterDeviceForbidden Forbidden

swagger:response registerDeviceForbidden
*/
type RegisterDeviceForbidden struct {
}

// NewRegisterDeviceForbidden creates RegisterDeviceForbidden with default headers values
func NewRegisterDeviceForbidden() *RegisterDeviceForbidden {

	return &RegisterDeviceForbidden{}
}

// WriteResponse to the client
func (o *RegisterDeviceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*RegisterDeviceDefault Error

swagger:response registerDeviceDefault
*/
type RegisterDeviceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterDeviceDefault creates RegisterDeviceDefault with default headers values
func NewRegisterDeviceDefault(code int) *RegisterDeviceDefault {
	if code <= 0 {
		code = 500
	}

	return &RegisterDeviceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the register device default response
func (o *RegisterDeviceDefault) WithStatusCode(code int) *RegisterDeviceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the register device default response
func (o *RegisterDeviceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the register device default response
func (o *RegisterDeviceDefault) WithPayload(payload *models.Error) *RegisterDeviceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register device default response
func (o *RegisterDeviceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterDeviceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}