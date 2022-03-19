package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/cecepsprd/grpc-gateway-boilerplate/api/proto"
	"github.com/cecepsprd/grpc-gateway-boilerplate/utils/logger"
	"google.golang.org/grpc"
)

// RunServer runs gRPC service to publish Catalog service
func RunServer(ctx context.Context, srv pb.UserServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(server, srv)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
