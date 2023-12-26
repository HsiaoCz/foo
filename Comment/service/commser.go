package service

import (
	"github.com/HsiaoCz/foo/Comment/data"
	"github.com/HsiaoCz/foo/Comment/pb/pv1"
)

type CommService struct {
	ur data.CommentRepo
	pv1.UnimplementedFooServer
}

func NewCommService(ur data.CommentRepo) *CommService {
	return &CommService{ur: ur}
}
