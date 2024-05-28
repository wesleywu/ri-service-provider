package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type DeleteMultiLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewDeleteMultiLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *DeleteMultiLogic {
	return &DeleteMultiLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *DeleteMultiLogic) DeleteMulti(ctx context.Context, req *p.VideoCollectionDeleteMultiReq) (*p.VideoCollectionDeleteMultiRes, error) {
	return s.repo.DeleteMulti(ctx, req)
}
