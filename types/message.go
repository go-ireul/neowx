package types

import (
	"encoding/xml"
	"strings"
)

const (
	// Text message type 'text'
	Text = "text"
	// Image message type 'image'
	Image = "image"
	// Voice message type 'voice'
	Voice = "voice"
	// Video message type 'viceo'
	Video = "video"
	// ShortVideo message type 'shortvideo'
	ShortVideo = "shortvideo"
	// Link message type 'link'
	Link = "link"
	// Location message type 'location', or event type 'LOCATION' ignore case
	Location = "location"
	// Event message type 'event'
	Event = "event"
	// Subscribe event 'subscribe'
	Subscribe = "subscribe"
	// Unsubscribe event 'unsubscribe'
	Unsubscribe = "unsubscribe"
	// QRScenePrefix prefix for 'subscribe' with QR scan
	QRScenePrefix = "qrscene_"
	// Scan event 'SCAN', ignore case
	Scan = "scan"
	// Click event 'CLICK', ignore case
	Click = "click"
	// View event 'VIEW', ignore case
	View = "view"
)

// Message incomming message from wechat server
type Message struct {
	XMLName xml.Name `xml:"xml"`
	// Basic Fields, not empty
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	// Message
	MsgID string `xml:"MsgId" json:"MsgId,omitempty"`
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
