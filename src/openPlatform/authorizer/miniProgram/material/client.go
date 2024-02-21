package material

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	response4 "github.com/stuartjing/PowerWeChat/v3/src/work/media/response"
	"net/http"
)

type Client struct {
	*kernel.BaseClient

	AllowTypes []string
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
		AllowTypes: []string{"image", "voice", "video", "thumb", "news_image"},
	}, nil
}

// 获取永久素材图片
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) Get(ctx context.Context, mediaID string) (*http.Response, error) {

	header := &response4.ResponseHeaderMedia{}
	response, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/material/get_material", http.MethodPost, &object.HashMap{
		"form_params": &object.HashMap{
			"media_id": mediaID,
		},
	}, header, nil)

	return response, err
}
