package service

import (
	"context"

	"github.com/HsiaoCz/foo/User/pb/pv1"
)

type UserCase struct {
	pv1.UnimplementedFooServer
}

func (u *UserCase) UserSignup(ctx context.Context, in *pv1.SignupRequest) (*pv1.SignupResponse, error) {
	return &pv1.SignupResponse{}, nil
}

func (u *UserCase) UserLogin(ctx context.Context, in *pv1.LoginRequest) (*pv1.LoginResponse, error) {
	return &pv1.LoginResponse{}, nil
}

func NewUserCase() *UserCase {
	return &UserCase{}
}
