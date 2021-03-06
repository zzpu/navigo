package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"ucenter/internal/biz"

	pb "ucenter/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	uc  *biz.MemberMgr
	log *log.Helper
}

func NewAuthService(uc *biz.MemberMgr, logger log.Logger) *AuthService {
	return &AuthService{uc: uc, log: log.NewHelper("service/auth", logger)}
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
func (s *AuthService) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoReply, error) {
	return &pb.GetInfoReply{}, nil
}
func (s *AuthService) UpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest) (*pb.UpdateInfoReply, error) {
	return &pb.UpdateInfoReply{}, nil
}
