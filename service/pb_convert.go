package service

import (
	pb "github.com/cecepsprd/grpc-gateway-boilerplate/api/proto"
	"github.com/cecepsprd/grpc-gateway-boilerplate/model"
)

func getUsersProto(users []model.User) *pb.Users {
	var resp pb.Users
	for _, user := range users {
		resp.Users = append(resp.Users, &pb.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			Phone:     user.Phone,
			Address:   user.Address,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		})
	}

	return &resp
}
