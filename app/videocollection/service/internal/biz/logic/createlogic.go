package logic

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type CreateLogic struct {
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

func NewCreateLogic(repo biz.VideoCollectionRepo, helper *log.Helper) *CreateLogic {
	return &CreateLogic{
		repo:   repo,
		helper: helper,
	}
}

func (s *CreateLogic) Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error) {
	return s.repo.Create(ctx, req)
}
