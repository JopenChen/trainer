package answer

import (
	"net/http"

	"github.com/JopenChen/trainer/trainer/internal/logic/answer"
	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateAnswerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateAnswerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := answer.NewCreateAnswerLogic(r.Context(), svcCtx)
		resp, err := l.CreateAnswer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
