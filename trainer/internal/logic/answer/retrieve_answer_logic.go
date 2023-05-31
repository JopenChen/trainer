package answer

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetrieveAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveAnswerLogic {
	return &RetrieveAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RetrieveAnswerLogic) RetrieveAnswer(req *types.GeneralIDRequest) (resp *types.RetrieveAnswerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
