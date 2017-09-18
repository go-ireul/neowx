package types

import "encoding/xml"

// Message incomming message from wechat server
type Message struct {
	XMLName xml.Name `xml:"xml"`
	// Basic Fields, not empty
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	// Message
	MsgID string `xml:"MsgId,omitempty" json:"MsgId,omitempty"`
	// TextMessage
	Content string `xml:",omitempty" json:",omitempty"`
	// RichMessage
	MediaID string `xml:"MediaId,omitempty" json:"MediaId,omitempty"`
	// ImageMessage
	PicURL string `xml:"PicUrl,omitempty" json:"PicUrl,omitempty"`
	// VoiceMessage
	Format      string `xml:",omitempty" json:",omitempty"`
	Recognition string `xml:",omitempty" json:",omitempty"`
	// VideoMessage // ShortVideoMessage
	ThumbMediaID string `xml:"ThumbMediaId,omitempty" json:"ThumbMediaId,omitempty"`
	// LocationMessage
	LocationX string `xml:"Location_X,omitempty" json:"Location_X,omitempty"`
	LocationY string `xml:"Location_Y,omitempty" json:"Location_Y,omitempty"`
	Scale     string `xml:",omitempty" json:",omitempty"`
	Label     string `xml:",omitempty" json:",omitempty"`
	// LinkMessage
	Title       string `xml:",omitempty" json:",omitempty"`
	Description string `xml:",omitempty" json:",omitempty"`
	URL         string `xml:"Url,omitempty" json:"Url,omitempty"`
	// Event
	Event    string `xml:",omitempty" json:",omitempty"`
	EventKey string `xml:",omitempty" json:",omitempty"`
	// SubscribeEvent
	Ticket string `xml:",omitempty" json:",omitempty"`
	// LocationEvent
	Latitude  string `xml:",omitempty" json:",omitempty"`
	Longitude string `xml:",omitempty" json:",omitempty"`
	Precision string `xml:",omitempty" json:",omitempty"`
}
