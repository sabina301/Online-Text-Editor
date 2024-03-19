package configuration

import (
	"errors"
	"net"
	"os"
)

type GrpcConfiguration interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGrpcConfig() (GrpcConfiguration, error) {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")
	if len(host) == 0 || len(port) == 0 {
		return nil, errors.New("Grpc host or port not found :(")
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (grpcConfig *grpcConfig) Address() string {
	return net.JoinHostPort(grpcConfig.host, grpcConfig.port)
}
