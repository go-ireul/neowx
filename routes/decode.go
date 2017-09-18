package routes

import (
	"encoding/xml"

	"ireul.com/neowx/types"
	"ireul.com/web"
)

// DecodeWechatXML decode Wechat XML body
func DecodeWechatXML() interface{} {
	return func(ctx *web.Context) {
		m := types.Message{}
		bytes, err := ctx.Req.Body().Bytes()
		if err != nil {
			ctx.Error(400, "cannot read POST body")
			return
		}
		xml.Unmarshal(bytes, &m)
		ctx.Map(m)
	}
}
