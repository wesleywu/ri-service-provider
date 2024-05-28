package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type ListLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewListLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *ListLogic {
	return &ListLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *ListLogic) List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error) {
	return s.repo.List(ctx, req)
}
