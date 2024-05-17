package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muazwzxv/go-backend-masterclass/modules/users"
	"github.com/muazwzxv/go-backend-masterclass/pb"
	"github.com/muazwzxv/go-backend-masterclass/pkg/hash"
	"github.com/muazwzxv/go-backend-masterclass/pkg/worker"
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
		s.rpcServer.Log.Infof("[CreateUser] - failed create user: %w", err)
		return nil, status.Error(codes.Internal, "internal server")
	}

	payload := &worker.PayloadSendVerifyEmail{
		UserID: user.ID,
	}

  // TODO: worker and insert user should be in one transaction
	err = s.rpcServer.TaskDistributor.SendVerifyEmail(ctx, payload)
	if err != nil {
		s.rpcServer.Log.Infof("[worker] - failed to send verification email: %w", err)
		return nil, status.Errorf(codes.Internal, "internal server")
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
		return nil, status.Errorf(codes.Unauthenticated, "internal server")
	}

	userRes := &pb.User{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UserName:  user.UserName,
		CreatedAt: timestamppb.New(*user.CreatedAt),
	}

	if user.DeletedAt != nil {
		userRes.DeletedAt = timestamppb.New(*user.DeletedAt)
	}

	res := &pb.LoginUserResponse{
		User:        userRes,
		AccessToken: loginData.Token,
	}
	return res, nil
}
