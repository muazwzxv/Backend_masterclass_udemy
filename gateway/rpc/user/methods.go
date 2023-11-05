package user

import (
	"context"

	"github.com/muazwzxv/go-backend-masterclass/modules/users"
	"github.com/muazwzxv/go-backend-masterclass/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	arg := &users.CreateUser{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
	}

	user, err := h.m.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server")
	}

	res := &pb.CreateUserResponse{
		User: &pb.User{
			Id:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}

	return res, nil
}
