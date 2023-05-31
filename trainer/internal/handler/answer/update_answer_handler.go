package answer

import (
	"net/http"

	"github.com/JopenChen/trainer/trainer/internal/logic/answer"
	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateAnswerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAnswerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := answer.NewUpdateAnswerLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAnswer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
