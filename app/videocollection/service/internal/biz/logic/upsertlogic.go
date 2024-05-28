package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type UpsertLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewUpsertLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *UpsertLogic {
	return &UpsertLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *UpsertLogic) Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error) {
	return s.repo.Upsert(ctx, req)
}
