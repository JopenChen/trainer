package group

import (
	"net/http"

	"github.com/JopenChen/trainer/trainer/internal/logic/group"
	"github.com/JopenChen/trainer/trainer/internal/svc"
	"github.com/JopenChen/trainer/trainer/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RetrieveGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GeneralIDRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewRetrieveGroupLogic(r.Context(), svcCtx)
		resp, err := l.RetrieveGroup(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
