package media

import (
	"net/http"

	mediaSvc "stack-bm/internal/service/mkt/media"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// OAuthCallbackHandler 公开回调入口（媒体平台授权回调，无需鉴权）
type OAuthCallbackHandler struct {
	auth *mediaSvc.ChannelAuthService
}

func NewOAuthCallbackHandler() *OAuthCallbackHandler {
	return &OAuthCallbackHandler{auth: mediaSvc.NewChannelAuthService()}
}

// Callback 接收媒体平台回调 code/state，完成授权
func (h *OAuthCallbackHandler) Callback(c *gin.Context) {
	channel := c.Param("channel")
	params := map[string]string{}
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	if err := h.auth.FinishOauth(channel, params); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}
