package question

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveQuestionLogic {
	return &RetrieveQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveQuestionLogic) RetrieveQuestion(req *types.GeneralIDRequest) (resp *types.RetrieveQuestionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
