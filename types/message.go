package types

import "encoding/xml"

// Message incomming message from wechat server
type Message struct {
	XMLName xml.Name `xml:"xml"`
	// Basic Fields, not empty
	ToUserName   string `xml:",cdata"`
	FromUserName string `xml:",cdata"`
	CreateTime   string
	MsgType      string `xml:",cdata"`
	// Message
	MsgID string `xml:"MsgId,omitempty" json:"MsgId,omitempty"`
	// TextMessage
	Content string `xml:",cdata,omitempty" json:",omitempty"`
	// RichMessage
	MediaID string `xml:"MediaId,cdata,omitempty" json:"MediaId,omitempty"`
	// ImageMessage
	PicURL string `xml:"PicUrl,cdata,omitempty" json:"PicUrl,omitempty"`
	// VoiceMessage
	Format      string `xml:",cdata,omitempty" json:",omitempty"`
	Recognition string `xml:",cdata,omitempty" json:",omitempty"`
	// VideoMessage // ShortVideoMessage
	ThumbMediaID string `xml:"ThumbMediaId,cdata,omitempty" json:"ThumbMediaId,omitempty"`
	// LocationMessage
	LocationX string `xml:"Location_X,omitempty" json:"Location_X,omitempty"`
	LocationY string `xml:"Location_Y,omitempty" json:"Location_Y,omitempty"`
	Scale     string `xml:",omitempty" json:",omitempty"`
	Label     string `xml:",cdata,omitempty" json:",omitempty"`
	// LinkMessage
	Title       string `xml:",cdata,omitempty" json:",omitempty"`
	Description string `xml:",cdata,omitempty" json:",omitempty"`
	URL         string `xml:"Url,cdata,omitempty" json:"Url,omitempty"`
	// Event
	Event    string `xml:",cdata,omitempty" json:",omitempty"`
	EventKey string `xml:",cdata,omitempty" json:",omitempty"`
	// SubscribeEvent
	Ticket string `xml:",cdata,omitempty" json:",omitempty"`
	// LocationEvent
	Latitude  string `xml:",omitempty" json:",omitempty"`
	Longitude string `xml:",omitempty" json:",omitempty"`
	Precision string `xml:",omitempty" json:",omitempty"`
}
