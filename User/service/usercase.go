package service

import (
	"context"

	"github.com/HsiaoCz/foo/User/pb/upv1"
)

type UserCase struct {
	upv1.UnimplementedUserPartServer
}

func (u *UserCase) UserSignup(ctx context.Context, in *upv1.SignupRequest) (*upv1.SignupResponse, error) {
	return &upv1.SignupResponse{}, nil
}

func (u *UserCase) UserLogin(ctx context.Context, in *upv1.LoginRequest) (*upv1.LoginResponse, error) {
	return &upv1.LoginResponse{}, nil
}

func NewUserCase() *UserCase {
	return &UserCase{}
}
