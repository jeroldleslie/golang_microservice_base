package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	io "go-microservice-base/auth/pkg/io"
	service "go-microservice-base/auth/pkg/service"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	UserDetails io.UserDetails `json:"user_details"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := s.Login(ctx, req.UserDetails)
		return LoginResponse{
			Err:   err,
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// LogoutRequest collects the request parameters for the Logout method.
type LogoutRequest struct{}

// LogoutResponse collects the response parameters for the Logout method.
type LogoutResponse struct {
	Err error `json:"err"`
}

// MakeLogoutEndpoint returns an endpoint that invokes Logout on the service.
func MakeLogoutEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := s.Logout(ctx)
		return LogoutResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r LogoutResponse) Failed() error {
	return r.Err
}

// ValidateTokenRequest collects the request parameters for the ValidateToken method.
type ValidateTokenRequest struct{}

// ValidateTokenResponse collects the response parameters for the ValidateToken method.
type ValidateTokenResponse struct {
	Session io.Session `json:"session"`
	Err     error      `json:"err"`
}

// MakeValidateTokenEndpoint returns an endpoint that invokes ValidateToken on the service.
func MakeValidateTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		session, err := s.ValidateToken(ctx)
		return ValidateTokenResponse{
			Err:     err,
			Session: session,
		}, nil
	}
}

// Failed implements Failer.
func (r ValidateTokenResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, userDetails io.UserDetails) (token string, err error) {
	request := LoginRequest{UserDetails: userDetails}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Token, response.(LoginResponse).Err
}

// Logout implements Service. Primarily useful in a client.
func (e Endpoints) Logout(ctx context.Context) (err error) {
	request := LogoutRequest{}
	response, err := e.LogoutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LogoutResponse).Err
}

// ValidateToken implements Service. Primarily useful in a client.
func (e Endpoints) ValidateToken(ctx context.Context) (session io.Session, err error) {
	request := ValidateTokenRequest{}
	response, err := e.ValidateTokenEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ValidateTokenResponse).Session, response.(ValidateTokenResponse).Err
}
