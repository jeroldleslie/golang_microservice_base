package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	io "go-microservice-base/auth/pkg/io"
)

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Login(ctx context.Context, userDetails io.UserDetails) (token string, err error) {
	defer func() {
		l.logger.Log("method", "Login", "userDetails", userDetails, "token", token, "err", err)
	}()
	return l.next.Login(ctx, userDetails)
}
func (l loggingMiddleware) Logout(ctx context.Context) (err error) {
	defer func() {
		l.logger.Log("method", "Logout", "err", err)
	}()
	return l.next.Logout(ctx)
}
func (l loggingMiddleware) ValidateToken(ctx context.Context) (session io.Session, err error) {
	defer func() {
		l.logger.Log("method", "ValidateToken", "session", session, "err", err)
	}()
	return l.next.ValidateToken(ctx)
}
