package service

import "github.com/HsiaoCz/foo/Content/pb/pv1"

type ContentService struct {
	pv1.UnimplementedFooServer
}

func NewContentService() *ContentService {
	return &ContentService{}
}
