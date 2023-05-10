package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-blog/app/core/service/internal/biz"

	v1 "kratos-blog/gen/api/go/user/service/v1"
)

type UserService struct {
	v1.UnimplementedUserServiceServer

	uc  *biz.UserUseCase
	tuc *biz.UserTokenUseCase
	log *log.Helper
}

func NewUserService(logger log.Logger, uc *biz.UserUseCase, tuc *biz.UserTokenUseCase) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/core-service"))
	return &UserService{
		log: l,
		uc:  uc,
		tuc: tuc,
	}
}

// Login 登陆
func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	_, err := s.uc.VerifyPassword(ctx, req)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	user, err := s.uc.GetUserByUserName(ctx, req.GetUserName())
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	token, err := s.tuc.GenerateToken(ctx, user)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		Token:    token,
		Id:       user.GetId(),
		UserName: user.GetUserName(),
	}, nil
}

// Logout 登出
func (s *UserService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.tuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// Register 注册
func (s *UserService) Register(ctx context.Context, req *v1.CreateUserRequest) (*v1.RegisterResponse, error) {
	user, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	token, err := s.tuc.GenerateToken(ctx, user)
	if err != nil {
		return &v1.RegisterResponse{}, err
	}

	return &v1.RegisterResponse{
		Id:       user.GetId(),
		UserName: user.GetUserName(),
		Token:    token,
	}, nil
}

func (s *UserService) GetMe(ctx context.Context, req *v1.GetMeRequest) (*v1.User, error) {
	return s.uc.Get(ctx, req.GetId())
}
