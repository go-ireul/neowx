package types

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
