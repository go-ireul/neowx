package types

import (
	"encoding/xml"
	"strings"
)

// WxReq incomming message from wechat server
type WxReq struct {
	XMLName xml.Name `xml:"xml"`
	// Basic Fields, not empty
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	// WxReq
	MsgID string `xml:"MsgId" json:"MsgId,omitempty"`
	// TextWxReq
	Content string `json:",omitempty"`
	// RichWxReq
	MediaID string `xml:"MediaId" json:"MediaId,omitempty"`
	// ImageWxReq
	PicURL string `xml:"PicUrl" json:"PicUrl,omitempty"`
	// VoiceWxReq
	Format      string `json:",omitempty"`
	Recognition string `json:",omitempty"`
	// VideoWxReq // ShortVideoWxReq
	ThumbMediaID string `xml:"ThumbMediaId" json:"ThumbMediaId,omitempty"`
	// LocationWxReq
	LocationX string `xml:"Location_X" json:"Location_X,omitempty"`
	LocationY string `xml:"Location_Y" json:"Location_Y,omitempty"`
	Scale     string `json:",omitempty"`
	Label     string `json:",omitempty"`
	// LinkWxReq
	Title       string `json:",omitempty"`
	Description string `json:",omitempty"`
	URL         string `xml:"Url" json:"Url,omitempty"`
	// Event
	Event    string `json:",omitempty"`
	EventKey string `json:",omitempty"`
	// SubscribeEvent
	Ticket string `json:",omitempty"`
	// LocationEvent
	Latitude  string `json:",omitempty"`
	Longitude string `json:",omitempty"`
	Precision string `json:",omitempty"`
}

// Is returns if this message is given type
func (m WxReq) Is(t string) bool {
	return strings.EqualFold(m.MsgType, t)
}

// IsEvent returns if this message is 'event' and is given event type
func (m WxReq) IsEvent(e string) bool {
	return m.Is(Event) && strings.EqualFold(m.Event, e)
}

// QRCode extract QRCode from both subscribe event or a single scan event
func (m WxReq) QRCode() (code string) {
	if m.Is(Event) {
		if m.IsEvent(Subscribe) {
			if strings.HasPrefix(m.EventKey, QRScenePrefix) {
				code = m.EventKey[len(QRScenePrefix):]
				return
			}
		}
		if m.IsEvent(Scan) {
			code = m.EventKey
			return
		}
	}
	return
}
