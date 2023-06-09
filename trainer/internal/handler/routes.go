// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	answer "github.com/JopenChen/trainer/trainer/internal/handler/answer"
	group "github.com/JopenChen/trainer/trainer/internal/handler/group"
	question "github.com/JopenChen/trainer/trainer/internal/handler/question"
	"github.com/JopenChen/trainer/trainer/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/group",
				Handler: group.CreateGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/group/:id",
				Handler: group.UpdateGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/:id",
				Handler: group.RetrieveGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/group/:id",
				Handler: group.RemoveGroupHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/question",
				Handler: question.CreateQuestionHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/question/:id",
				Handler: question.UpdateQuestionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/question/:id",
				Handler: question.RetrieveQuestionHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/question/:id",
				Handler: question.RemoveQuestionHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/answer",
				Handler: answer.CreateAnswerHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/answer/:id",
				Handler: answer.UpdateAnswerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/answer/:id",
				Handler: answer.RetrieveAnswerHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/answer/:id",
				Handler: answer.RemoveAnswerHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
