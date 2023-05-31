package question

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveQuestionLogic {
	return &RemoveQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveQuestionLogic) RemoveQuestion(req *types.GeneralIDRequest) (resp *types.GeneralResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
