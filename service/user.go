package service

import (
	"context"
	"time"

	pb "github.com/cecepsprd/grpc-gateway-boilerplate/api/proto"
	"github.com/cecepsprd/grpc-gateway-boilerplate/repository"
	"github.com/cecepsprd/grpc-gateway-boilerplate/utils/logger"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserService interface {
	Read(ctx context.Context, empty *empty.Empty) (users *pb.Users, err error)
}

type userService struct {
	repo           repository.UserRepository
	contextTimeout time.Duration
}

func NewUserService(urepo repository.UserRepository, timeout time.Duration) UserService {
	return &userService{
		repo:           urepo,
		contextTimeout: timeout,
	}
}

func (s *userService) Read(ctx context.Context, empty *empty.Empty) (*pb.Users, error) {
	users, err := s.repo.Read(ctx)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}

	return getUsersProto(users), nil
}
