package group

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveGroupLogic {
	return &RetrieveGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveGroupLogic) RetrieveGroup(req *types.GeneralIDRequest) (resp *types.RetrieveGroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
