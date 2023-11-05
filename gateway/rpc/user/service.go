package user

import (
	"github.com/muazwzxv/go-backend-masterclass/modules/users"
	"github.com/muazwzxv/go-backend-masterclass/pb"
	"github.com/muazwzxv/go-backend-masterclass/pkg/rpcServer"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	rpcServer *rpcServer.Server
	m         users.IUsers
}

func NewUserServiceServer(rpc *rpcServer.Server, module users.IUsers) *UserService {
	return &UserService{
		rpcServer: rpc,
    m: module,
	}
}
