// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	models	"github.com/geneva_horodateur/models"
//	models "github.com/Magicking/rc-ge-ch-pdf/models"
)

// ListtimestampedOKCode is the HTTP code returned for type ListtimestampedOK
const ListtimestampedOKCode int = 200

/*ListtimestampedOK Liste des fichiers qui ont été horodaté


swagger:response listtimestampedOK
*/
type ListtimestampedOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ReceiptFile `json:"body,omitempty"`
}

// NewListtimestampedOK creates ListtimestampedOK with default headers values
func NewListtimestampedOK() *ListtimestampedOK {

	return &ListtimestampedOK{}
}

// WithPayload adds the payload to the listtimestamped o k response
func (o *ListtimestampedOK) WithPayload(payload []*models.ReceiptFile) *ListtimestampedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listtimestamped o k response
func (o *ListtimestampedOK) SetPayload(payload []*models.ReceiptFile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListtimestampedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.ReceiptFile, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ListtimestampedDefault Internal error

swagger:response listtimestampedDefault
*/
type ListtimestampedDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListtimestampedDefault creates ListtimestampedDefault with default headers values
func NewListtimestampedDefault(code int) *ListtimestampedDefault {
	if code <= 0 {
		code = 500
	}

	return &ListtimestampedDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the listtimestamped default response
func (o *ListtimestampedDefault) WithStatusCode(code int) *ListtimestampedDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the listtimestamped default response
func (o *ListtimestampedDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the listtimestamped default response
func (o *ListtimestampedDefault) WithPayload(payload *models.Error) *ListtimestampedDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listtimestamped default response
func (o *ListtimestampedDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListtimestampedDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
