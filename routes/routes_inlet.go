package routes

import (
	"fmt"
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
func InletPost(ctx *web.Context, m types.WxReq, cfg types.Config, rs []types.Rule) {
	for _, r := range rs {
		if r.Fn == nil || r.Fn(m, cfg) {
			if len(r.HTTPSync) > 0 {
				//TODO: HTTPSync and return
			} else if len(r.HTTPAsync) > 0 {
				//TODO: go HTTPSync
			}
			if len(r.Text) > 0 {
				resp := types.WxTextResp{
					ToUserName:   com.NewCDATA(m.FromUserName),
					FromUserName: com.NewCDATA(m.ToUserName),
					CreateTime:   fmt.Sprintf("%d", time.Now().Unix()),
					MsgType:      com.NewCDATA(types.Text),
					Content:      com.NewCDATA(r.Text),
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
