package grpc

import (
	"context"
	"go-microservice-base/notificator/pkg/endpoint"
	"go-microservice-base/notificator/pkg/grpc/pb"
	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeNotifyHandler creates the handler logic
func makeNotifyHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.NotifyEndpoint, decodeNotifyRequest, encodeNotifyResponse, options...)
}

// decodeNotifyResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeNotifyRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.NotifyRequest)
	return endpoint.NotifyRequest{Channel: req.Channel, Message: req.Message}, nil
}

// encodeNotifyResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeNotifyResponse(_ context.Context, r interface{}) (interface{}, error) {
	response := r.(endpoint.NotifyResponse)
	return &pb.NotifyReply{Response: response.Response}, nil
}
func (g *grpcServer) Notify(ctx context1.Context, req *pb.NotifyRequest) (*pb.NotifyReply, error) {
	_, rep, err := g.notify.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.NotifyReply), nil
}
