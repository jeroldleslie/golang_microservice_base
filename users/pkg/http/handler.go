package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"

	"github.com/jeroldleslie/golang_microservice_base/users/pkg/endpoint"
)

// makeCreateHandler creates the handler logic
func makeCreateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"POST", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(
			http.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)))
}

// decodeCreateResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	if err.Error() == "invalid id" {
		return http1.StatusBadRequest
	} else if err.Error() == "not found" {
		return http1.StatusNotFound
	}
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeGetByIdHandler creates the handler logic
func makeGetByIdHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/get-by-id/{id}").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(
			http.NewServer(endpoints.GetByIdEndpoint, decodeGetByIdRequest, encodeGetByIdResponse, options...)))
}

// decodeGetByIdResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIdRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIdRequest{
		Id: id,
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetByIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIdResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeHealthHandler creates the handler logic
func makeHealthHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/health").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(
			http.NewServer(endpoints.HealthEndpoint, decodeHealthRequest, encodeHealthResponse, options...)))
}

// decodeHealthResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeHealthRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.HealthRequest{}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeHealthResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeHealthResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
