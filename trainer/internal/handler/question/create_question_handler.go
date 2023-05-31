package question

import (
	"net/http"

	"github.com/JopenChen/trainer/trainer/internal/logic/question"
	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateQuestionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := question.NewCreateQuestionLogic(r.Context(), svcCtx)
		resp, err := l.CreateQuestion(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
