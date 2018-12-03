package endpoint

import (
	"context"
	service "go-microservice-base/notificator/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct{}

// HealthResponse collects the response parameters for the Health method.
type HealthResponse struct {
	Status bool `json:"status"`
}

// MakeHealthEndpoint returns an endpoint that invokes Health on the service.
func MakeHealthEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status := s.Health(ctx)
		return HealthResponse{Status: status}, nil
	}
}

// Health implements Service. Primarily useful in a client.
func (e Endpoints) Health(ctx context.Context) (status bool) {
	request := HealthRequest{}
	response, err := e.HealthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthResponse).Status
}

// NotifyRequest collects the request parameters for the Notify method.
type NotifyRequest struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

// NotifyResponse collects the response parameters for the Notify method.
type NotifyResponse struct {
	Response string `json:"response"`
	Err      error  `json:"err"`
}

// MakeNotifyEndpoint returns an endpoint that invokes Notify on the service.
func MakeNotifyEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NotifyRequest)
		response, err := s.Notify(ctx, req.Channel, req.Message)
		return NotifyResponse{
			Err:      err,
			Response: response,
		}, nil
	}
}

// Failed implements Failer.
func (r NotifyResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Notify implements Service. Primarily useful in a client.
func (e Endpoints) Notify(ctx context.Context, channel string, message string) (response string, err error) {
	request := NotifyRequest{
		Channel: channel,
		Message: message,
	}
	response0, err := e.NotifyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response0.(NotifyResponse).Response, response0.(NotifyResponse).Err
}
