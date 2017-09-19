package wx

import (
	"encoding/xml"
	"strings"

	"ireul.com/com"
)

// Message incomming message from wechat server
type Message struct {
	XMLName xml.Name `xml:"xml"`
	// Basic Fields, not empty
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	// Message
	MsgID int64 `xml:"MsgId" json:"MsgId,omitempty"`
	// TextMessage
	Content string `json:",omitempty"`
	// RichMessage
	MediaID string `xml:"MediaId" json:"MediaId,omitempty"`
	// ImageMessage
	PicURL string `xml:"PicUrl" json:"PicUrl,omitempty"`
	// VoiceMessage
	Format      string `json:",omitempty"`
	Recognition string `json:",omitempty"`
	// VideoMessage // ShortVideoMessage
	ThumbMediaID string `xml:"ThumbMediaId" json:"ThumbMediaId,omitempty"`
	// LocationMessage
	LocationX string `xml:"Location_X" json:"Location_X,omitempty"`
	LocationY string `xml:"Location_Y" json:"Location_Y,omitempty"`
	Scale     string `json:",omitempty"`
	Label     string `json:",omitempty"`
	// LinkMessage
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
func (m Message) Is(t string) bool {
	return strings.EqualFold(m.MsgType, t)
}

// IsEvent returns if this message is 'event' and is given event type
func (m Message) IsEvent(e string) bool {
	return m.Is(Event) && strings.EqualFold(m.Event, e)
}

// QRCode extract QRCode from both subscribe event or a single scan event
func (m Message) QRCode() (code string) {
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

// TextReply is a text response to Wechat server
type TextReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   com.CDATA
	FromUserName com.CDATA
	CreateTime   string
	MsgType      com.CDATA
	Content      com.CDATA
}
