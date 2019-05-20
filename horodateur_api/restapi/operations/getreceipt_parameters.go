// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetreceiptParams creates a new GetreceiptParams object
// no default values defined in spec.
func NewGetreceiptParams() GetreceiptParams {

	return GetreceiptParams{}
}

// GetreceiptParams contains all the bound params for the getreceipt operation
// typically these are obtained from a http.Request
//
// swagger:parameters getreceipt
type GetreceiptParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Le hash identifiant un fichier
	  Required: true
	  In: query
	*/
	Hash string
	/*Langue du reçu
	  In: query
	*/
	Lang *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetreceiptParams() beforehand.
func (o *GetreceiptParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qHash, qhkHash, _ := qs.GetOK("hash")
	if err := o.bindHash(qHash, qhkHash, route.Formats); err != nil {
		res = append(res, err)
	}

	qLang, qhkLang, _ := qs.GetOK("lang")
	if err := o.bindLang(qLang, qhkLang, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHash binds and validates parameter Hash from query.
func (o *GetreceiptParams) bindHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("hash", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("hash", "query", raw); err != nil {
		return err
	}

	o.Hash = raw

	return nil
}

// bindLang binds and validates parameter Lang from query.
func (o *GetreceiptParams) bindLang(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Lang = &raw

	return nil
}