package routes

import (
	"time"

	"ireul.com/com"
	"ireul.com/neowx/types"
	"ireul.com/web"
)

// InletGet GET to /inlet
func InletGet(ctx *web.Context) {
	ctx.PlainText(200, []byte(ctx.Query("echostr")))
}

// InletPost POST to /inlet
func InletPost(ctx *web.Context, m types.WxReq, cfg types.Config) {
	for _, rule := range cfg.Rules {
		ok, err := rule.Matches(m)
		if err != nil {
			ctx.Error(500, "failed to execute match: "+err.Error())
			return
		}
		if ok {
			if len(rule.HTTPSync) > 0 {
				//TODO: http sync and return
			}
			if len(rule.HTTPAsync) > 0 {
				//TODO: go http sync
			}
			if len(rule.Text) > 0 {
				resp := types.WxTextResp{
					FromUserName: com.NewCDATA(m.ToUserName),
					ToUserName:   com.NewCDATA(m.FromUserName),
					CreateTime:   com.ToStr(time.Now().Unix()),
					MsgType:      com.NewCDATA(types.Text),
					Content:      com.NewCDATA(rule.Text),
				}
				ctx.XML(200, resp)
			} else {
				ctx.PlainText(200, []byte("success"))
			}
			return
		}
	}
	ctx.PlainText(200, []byte("success"))
}
