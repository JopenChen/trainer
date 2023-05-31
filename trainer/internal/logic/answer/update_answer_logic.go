package answer

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnswerLogic {
	return &UpdateAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAnswerLogic) UpdateAnswer(req *types.UpdateAnswerRequest) (resp *types.GeneralResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
