package answer

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAnswerLogic {
	return &RemoveAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAnswerLogic) RemoveAnswer(req *types.GeneralIDRequest) (resp *types.GeneralResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
