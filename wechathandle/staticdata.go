package wechathandle

const (
	//msg types
	MsgTypeText       = "text"
	MsgTypeImage      = "image"
	MsgTypeVoice      = "voice"
	MsgTypeVideo      = "video"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLocation   = "location"
	MsgTypeLink       = "link"
	MsgTypeEvent      = "event"
	// event types
	EventSubscribe   = "subscribe"
	EventUnsubscribe = "unsubscribe"
	EventScan        = "SCAN"
	EventLocation    = "LOCATION"
	EventClick       = "CLICK"
	EventView        = "VIEW"
	// media types
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
	// button types
	ButtonTypeClick = "click"
	ButtonTypeView  = "view"
	// environment constants
	UrlPrefix      = "https://api.weixin.qq.com/cgi-bin/"
	MediaUrlPrefix = "http://file.api.weixin.qq.com/cgi-bin/media/"
	retryNum       = 3
)
