package logic

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TrainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTrainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrainerLogic {
	return &TrainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TrainerLogic) Trainer(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
