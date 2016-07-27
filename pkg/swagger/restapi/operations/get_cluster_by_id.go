package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetClusterByIDHandlerFunc turns a function with the right signature into a get cluster by Id handler
type GetClusterByIDHandlerFunc func(GetClusterByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClusterByIDHandlerFunc) Handle(params GetClusterByIDParams) middleware.Responder {
	return fn(params)
}

// GetClusterByIDHandler interface for that can handle valid get cluster by Id params
type GetClusterByIDHandler interface {
	Handle(GetClusterByIDParams) middleware.Responder
}

// NewGetClusterByID creates a new http.Handler for the get cluster by Id operation
func NewGetClusterByID(ctx *middleware.Context, handler GetClusterByIDHandler) *GetClusterByID {
	return &GetClusterByID{Context: ctx, Handler: handler}
}

/*GetClusterByID swagger:route GET /jing getClusterById

GetClusterByID get cluster by Id API

*/
type GetClusterByID struct {
	Context *middleware.Context
	Handler GetClusterByIDHandler
}

func (o *GetClusterByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetClusterByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
