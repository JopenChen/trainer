package answer

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAnswerLogic {
	return &CreateAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAnswerLogic) CreateAnswer(req *types.CreateAnswerRequest) (resp *types.GeneralResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
