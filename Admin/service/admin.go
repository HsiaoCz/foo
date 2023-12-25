package service

import "github.com/HsiaoCz/foo/Admin/pb/pv1"

type AdminSer struct {
	pv1.UnimplementedFooServer
}

func NewAdminSer() *AdminSer {
	return &AdminSer{}
}
