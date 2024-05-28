package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type DeleteLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewDeleteLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *DeleteLogic {
	return &DeleteLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *DeleteLogic) Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error) {
	return s.repo.Delete(ctx, req)
}
