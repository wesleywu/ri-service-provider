package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type GetLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewGetLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *GetLogic {
	return &GetLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *GetLogic) Get(ctx context.Context, req *p.VideoCollectionGetReq) (*p.VideoCollectionGetRes, error) {
	return s.repo.Get(ctx, req)
}
