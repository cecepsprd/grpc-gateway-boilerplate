package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cecepsprd/grpc-gateway-boilerplate/cmd/server/grpc"
	"github.com/cecepsprd/grpc-gateway-boilerplate/cmd/server/rest"
	"github.com/cecepsprd/grpc-gateway-boilerplate/config"
	"github.com/cecepsprd/grpc-gateway-boilerplate/repository"
	"github.com/cecepsprd/grpc-gateway-boilerplate/service"
	"github.com/cecepsprd/grpc-gateway-boilerplate/utils/logger"
)

func ServeGRPCAndHTTP() error {
	var ctx = context.Background()
	var cfg = config.NewConfig()

	// initialize logger
	if err := logger.Init(cfg.App.LogLevel, cfg.App.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	// connect to database
	conn, err := cfg.MysqlConnect()
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}
	defer conn.Close()

	// // Initialize repository
	userRepository := repository.NewUserRepository(conn)

	// // Initialize service
	userService := service.NewUserService(userRepository, time.Duration(cfg.App.ContextTimeout))

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.App.GRPCPort, cfg.App.HTTPPort)
	}()

	return grpc.RunServer(ctx, userService, cfg.App.GRPCPort)
}
