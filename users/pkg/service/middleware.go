package service

import (
	"context"
	io "fivekilometer/users/pkg/io"

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

func (l loggingMiddleware) Create(ctx context.Context, user_in io.User) (user_out io.User, error error) {
	defer func() {
		l.logger.Log("log","Service middleware logging now .................")
		l.logger.Log("method", "Create", "user_in", user_in, "user_out", user_out, "error", error)
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
