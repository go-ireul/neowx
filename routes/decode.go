package routes

import (
	"encoding/xml"
	"log"

	"ireul.com/neowx/types"
	"ireul.com/web"
)

// DecodeWechatXML decode Wechat XML body
func DecodeWechatXML() interface{} {
	return func(ctx *web.Context) {
		m := types.Message{}
		bytes, err := ctx.Req.Body().Bytes()
		if err != nil {
			log.Println(err.Error())
			ctx.Error(400, "Cannot read request body: "+err.Error())
			return
		}
		err = xml.Unmarshal(bytes, &m)
		if err != nil {
			log.Println(err.Error())
			ctx.Error(400, "Cannot unmarshal XML: "+err.Error())
			return
		}
		ctx.Map(m)
	}
}
