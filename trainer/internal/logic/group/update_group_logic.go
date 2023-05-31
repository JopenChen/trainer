package group

import (
	"context"

	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupLogic) UpdateGroup(req *types.UpdateGroupRequest) (resp *types.GeneralResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
