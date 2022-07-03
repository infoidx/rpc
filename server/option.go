package server

import "google.golang.org/grpc"

type Option struct {
	Net
	opts []grpc.ServerOption
}

type Net struct {
	NetWork string // tcp
	Addr    string // :50000
}

type GrpcServerOption func() grpc.ServerOption
type OptionConfigure func(option *Option) error

func (o *Option) ApplyGrpcServerOption(grpcOpts ...GrpcServerOption) {
	for _, fn := range grpcOpts {
		o.opts = append(o.opts, fn())
	}
}
