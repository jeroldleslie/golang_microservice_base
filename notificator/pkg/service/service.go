package service

import (
	"context"
)

// NotificatorService describes the service.
type NotificatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Health(ctx context.Context) (status bool)
	Notify(ctx context.Context, channel string, message string) (response string, err error)
}

type basicNotificatorService struct{}

func (b *basicNotificatorService) Health(ctx context.Context) (status bool) {
	status = true
	return status
}

// NewBasicNotificatorService returns a naive, stateless implementation of NotificatorService.
func NewBasicNotificatorService() NotificatorService {
	return &basicNotificatorService{}
}

// New returns a NotificatorService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificatorService {
	var svc NotificatorService = NewBasicNotificatorService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicNotificatorService) Notify(ctx context.Context, channel string, message string) (response string, err error) {
	// TODO dummy business logic ---- implement the business logic of Notify
	response = "successfully notified"
	return response, err
}
