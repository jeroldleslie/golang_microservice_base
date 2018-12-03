package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(NotificatorService) NotificatorService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificatorService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NotificatorService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificatorService) NotificatorService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Health(ctx context.Context) (status bool) {
	defer func() {
		l.logger.Log("method", "Health", "status", status)
	}()
	return l.next.Health(ctx)
}

func (l loggingMiddleware) Notify(ctx context.Context, channel string, message string) (response string, err error) {
	defer func() {
		l.logger.Log("method", "Notify", "channel", channel, "message", message, "response", response, "err", err)
	}()
	return l.next.Notify(ctx, channel, message)
}
