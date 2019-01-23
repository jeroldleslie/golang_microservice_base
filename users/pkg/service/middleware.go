package service

import (
	"context"
	io "go-microservice-base/users/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UsersService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, user_in io.User) (error error) {
	defer func() {
		l.logger.Log("method", "Create", "user_in", user_in, "error", error)
	}()
	return l.next.Create(ctx, user_in)
}

func (l loggingMiddleware) GetById(ctx context.Context, id string) (u io.User, error error) {
	defer func() {
		l.logger.Log("method", "GetById", "id", id, "u", u, "error", error)
	}()
	return l.next.GetById(ctx, id)
}

func (l loggingMiddleware) Health(ctx context.Context) (status bool) {
	defer func() {
		l.logger.Log("method", "Health", "status", status)
	}()
	return l.next.Health(ctx)
}

func (l loggingMiddleware) Login(ctx context.Context, auth io.Authentication) (token string, error error) {
	defer func() {
		l.logger.Log("method", "Login", "auth", auth, "token", token, "error", error)
	}()
	return l.next.Login(ctx, auth)
}
