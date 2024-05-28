package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type UpdateLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewUpdateLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *UpdateLogic {
	return &UpdateLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *UpdateLogic) Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error) {
	return s.repo.Update(ctx, req)
}
