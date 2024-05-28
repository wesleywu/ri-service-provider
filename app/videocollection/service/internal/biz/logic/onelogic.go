package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type OneLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewOneLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *OneLogic {
	return &OneLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *OneLogic) One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error) {
	return s.repo.One(ctx, req)
}
