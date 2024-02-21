package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
)

type NewsItem struct {
	*Message
}

func NewNewsItem(data *object.HashMap) *NewsItem {
	m := &NewsItem{
		NewMessage(&power.HashMap{"items": data}),
	}

	m.Type = "news"
	m.Properties = []string{"title", "description", "url", "image"}
	m.OverrideToXmlArray()

	return m
}

func (msg *NewsItem) ToJsonArray() *object.HashMap {
	return &object.HashMap{
		"Title":       msg.Get("title", nil),
		"Description": msg.Get("description", nil),
		"Url":         msg.Get("url", nil),
		"PicUrl":      msg.Get("image", nil),
	}
}

// Override ToXmlArray
func (msg *NewsItem) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		return &object.HashMap{
			"Title":       msg.Get("title", nil),
			"Description": msg.Get("description", nil),
			"Url":         msg.Get("url", nil),
			"PicUrl":      msg.Get("image", nil),
		}
	}
}
