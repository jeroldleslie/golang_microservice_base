package service

import (
	"context"
	"go-microservice-base/auth/pkg/io"
)

// AuthService describes the service.
type AuthService interface {
	// Add your methods here
	Login(ctx context.Context, userDetails io.UserDetails) (token string, err error)
	Logout(ctx context.Context) (err error)
	ValidateToken(ctx context.Context) (session io.Session, err error)
}

type basicAuthService struct{}

func (b *basicAuthService) Login(ctx context.Context, userDetails io.UserDetails) (token string, err error) {
	// TODO implement the business logic of Login
	return token, err
}
func (b *basicAuthService) Logout(ctx context.Context) (err error) {
	// TODO implement the business logic of Logout
	return err
}
func (b *basicAuthService) ValidateToken(ctx context.Context) (session io.Session, err error) {
	// TODO implement the business logic of ValidateToken
	return session, err
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService() AuthService {
	return &basicAuthService{}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
