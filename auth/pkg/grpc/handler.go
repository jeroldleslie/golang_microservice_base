package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "go-microservice-base/auth/pkg/endpoint"
	pb "go-microservice-base/auth/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeLoginHandler creates the handler logic
func makeLoginHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)
}

// decodeLoginResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Decoder is not impelemented")
}

// encodeLoginResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Encoder is not impelemented")
}
func (g *grpcServer) Login(ctx context1.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	_, rep, err := g.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginReply), nil
}

// makeLogoutHandler creates the handler logic
func makeLogoutHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LogoutEndpoint, decodeLogoutRequest, encodeLogoutResponse, options...)
}

// decodeLogoutResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeLogoutRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Decoder is not impelemented")
}

// encodeLogoutResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeLogoutResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Encoder is not impelemented")
}
func (g *grpcServer) Logout(ctx context1.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	_, rep, err := g.logout.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LogoutReply), nil
}

// makeValidateTokenHandler creates the handler logic
func makeValidateTokenHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ValidateTokenEndpoint, decodeValidateTokenRequest, encodeValidateTokenResponse, options...)
}

// decodeValidateTokenResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeValidateTokenRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Decoder is not impelemented")
}

// encodeValidateTokenResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeValidateTokenResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Encoder is not impelemented")
}
func (g *grpcServer) ValidateToken(ctx context1.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenReply, error) {
	_, rep, err := g.validateToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ValidateTokenReply), nil
}
