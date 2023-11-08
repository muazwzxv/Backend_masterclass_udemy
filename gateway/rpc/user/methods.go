package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muazwzxv/go-backend-masterclass/modules/users"
	"github.com/muazwzxv/go-backend-masterclass/pb"
	"github.com/muazwzxv/go-backend-masterclass/pkg/hash"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	arg := &users.CreateUser{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		UserName:  req.GetUserName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
	}

	user, err := s.m.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server")
	}

	res := &pb.CreateUserResponse{
		User: &pb.User{
			Id:        user.ID,
			FirstName: user.FirstName,
			UserName:  user.UserName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}

	return res, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := s.m.FindUserByUserName(ctx, req.GetUserName())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "not found")
		}
		return nil, status.Errorf(codes.Internal, "internal server")
	}

	err = hash.CheckPassword(req.GetPassword(), user.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "internal server")
	}

	loginData, err := s.m.LoginUser(ctx, &users.LoginUserRequest{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
	})
  if err != nil {
    // TODO - proper error handling
		return nil, status.Errorf(codes.Unauthenticated, "internal server")
  }

  res := &pb.LoginUserResponse{
    User: &pb.User{
      Id: user.ID,
      FirstName: user.FirstName,
      LastName: user.LastName,
      Email: user.Email,
      UserName: user.UserName,
      CreatedAt: timestamppb.New(*user.CreatedAt),
      DeletedAt: timestamppb.New(*user.DeletedAt),
    },
    AccessToken: loginData.Token,
  }
  return res, nil
}
