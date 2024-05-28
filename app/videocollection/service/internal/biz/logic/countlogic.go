package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type CountLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewCountLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *CountLogic {
	return &CountLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *CountLogic) Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error) {
	return s.repo.Count(ctx, req)
}
