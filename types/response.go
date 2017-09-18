package types

import (
	"encoding/xml"

	"ireul.com/com"
)

// WxTextResp is a text response to Wechat server
type WxTextResp struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   com.CDATA
	FromUserName com.CDATA
	CreateTime   string
	MsgType      com.CDATA
	Content      com.CDATA
}
