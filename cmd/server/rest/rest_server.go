package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/cecepsprd/grpc-gateway-boilerplate/api/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		// Disable omitempty in response
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		// Modify http status code in response
		// runtime.WithForwardResponseOption(utils.HTTPResponseModifier),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalln("failed to start HTTP gateway")
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", httpPort),
		// add handler with middleware
		// Handler: middleware.AddRequestID(
		// 	middleware.AddLogger(logger.Log, mux)),
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		log.Println("starting HTTP/REST gateway...")
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	// Block until a signal is received.
	<-quit

	log.Println("server shutdown of 5 second.")

	// gracefully shutdown the server, waiting max 5 seconds for current operations to complete
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}
