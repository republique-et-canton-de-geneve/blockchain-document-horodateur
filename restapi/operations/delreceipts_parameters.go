// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDelreceiptsParams creates a new DelreceiptsParams object
// no default values defined in spec.
func NewDelreceiptsParams() DelreceiptsParams {

	return DelreceiptsParams{}
}

// DelreceiptsParams contains all the bound params for the delreceipts operation
// typically these are obtained from a http.Request
//
// swagger:parameters delreceipts
type DelreceiptsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Liste des hash à supprimer
	  Required: true
	  In: body
	*/
	Hashes []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDelreceiptsParams() beforehand.
func (o *DelreceiptsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("hashes", "body"))
			} else {
				res = append(res, errors.NewParseError("hashes", "body", "", err))
			}
		} else {

			// validate inline body array
			o.Hashes = body
			if err := o.validateHashesBody(route.Formats); err != nil {
				res = append(res, err)
			}
		}
	} else {
		res = append(res, errors.Required("hashes", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// validateHashesBody validates an inline body parameter
func (o *DelreceiptsParams) validateHashesBody(formats strfmt.Registry) error {

	return nil
}