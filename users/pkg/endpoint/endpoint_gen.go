// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "go-microservice-base/users/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	HealthEndpoint  endpoint.Endpoint
	CreateEndpoint  endpoint.Endpoint
	GetByIdEndpoint endpoint.Endpoint
	LoginEndpoint   endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.UsersService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateEndpoint:  MakeCreateEndpoint(s),
		GetByIdEndpoint: MakeGetByIdEndpoint(s),
		HealthEndpoint:  MakeHealthEndpoint(s),
		LoginEndpoint:   MakeLoginEndpoint(s),
	}
	for _, m := range mdw["Health"] {
		eps.HealthEndpoint = m(eps.HealthEndpoint)
	}
	for _, m := range mdw["Create"] {
		eps.CreateEndpoint = m(eps.CreateEndpoint)
	}
	for _, m := range mdw["GetById"] {
		eps.GetByIdEndpoint = m(eps.GetByIdEndpoint)
	}
	for _, m := range mdw["Login"] {
		eps.LoginEndpoint = m(eps.LoginEndpoint)
	}
	return eps
}
