package group

import (
	"context"
	"errors"
	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"
	"github.com/JopenChen/trainer/trainer/model"
	"github.com/JopenChen/trainer/trainer/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupRequest) (resp *types.GeneralResponse, err error) {
	resp = new(types.GeneralResponse)

	conditionMap := make(map[string]interface{})
	conditionMap["eq_name"] = req.Name

	group, err := l.svcCtx.GroupModel.GetByCondition(l.ctx, conditionMap)
	if err != nil && err != model.ErrNotFound {
		resp.Msg = err.Error()
		return
	}
	if group != nil {
		err = errors.New("record exists")
		resp.Msg = err.Error()
		return
	}

	_, err = l.svcCtx.GroupModel.Insert(l.ctx, &model.Group{
		Id:   util.GenerateID(),
		Name: req.Name,
	})
	if err != nil {
		resp.Msg = err.Error()
		return
	}

	resp.Msg = "OK"

	return
}
