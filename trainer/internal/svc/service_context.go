package svc

import (
	"github.com/JopenChen/trainer/trainer/internal/config"
	"github.com/JopenChen/trainer/trainer/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	GroupModel    model.GroupModel
	QuestionModel model.QuestionModel
	AnswerModel   model.AnswerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:        c,
		GroupModel:    model.NewGroupModel(conn),
		QuestionModel: model.NewQuestionModel(conn),
		AnswerModel:   model.NewAnswerModel(conn),
	}
}
