package http

import (
	"context"
	"encoding/json"
	endpoint "fivekilometer/notificator/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// makeHealthHandler creates the handler logic
func makeHealthHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Methods("GET").Path("/health").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(
			http1.NewServer(endpoints.HealthEndpoint, decodeHealthRequest, encodeHealthResponse, options...)))
}


// decodeHealthResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.HealthRequest{}
	return req, nil
}

// encodeHealthResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeHealthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
