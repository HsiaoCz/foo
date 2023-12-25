package service

import "github.com/HsiaoCz/foo/Serach/pb/pv1"

type SerachService struct {
	pv1.UnimplementedFooServer
}

func NewSerachService() *SerachService {
	return &SerachService{}
}
