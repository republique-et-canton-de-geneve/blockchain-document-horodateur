// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"errors"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DelreceiptsHandlerFunc turns a function with the right signature into a delreceipts handler
type DelreceiptsHandlerFunc func(DelreceiptsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DelreceiptsHandlerFunc) Handle(params DelreceiptsParams) middleware.Responder {
	return fn(params)
}

// DelreceiptsHandler interface for that can handle valid delreceipts params
type DelreceiptsHandler interface {
	Handle(DelreceiptsParams) middleware.Responder
}

// NewDelreceipts creates a new http.Handler for the delreceipts operation
func NewDelreceipts(ctx *middleware.Context, handler DelreceiptsHandler) *Delreceipts {
	return &Delreceipts{Context: ctx, Handler: handler}
}

/*Delreceipts swagger:route POST /recu delreceipts

Supprime les reçus de la base de donnée

Supprimer les reçus de la base de donnée


*/
type Delreceipts struct {
	Context *middleware.Context
	Handler DelreceiptsHandler
}

func (o *Delreceipts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDelreceiptsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, errors.New("Wrong params"))
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
