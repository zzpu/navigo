
package service

import(
	"context"

	pb "ucenter/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() pb.AuthServer {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *AuthService) Info(ctx context.Context, req *pb.InfoRequest) (*pb.InfoReply, error) {
	return &pb.InfoReply{}, nil
}
