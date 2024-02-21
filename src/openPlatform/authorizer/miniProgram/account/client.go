package account

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	response2 "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/aggregate/account"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/account/response"
)

type Client struct {
	*account.Client
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	accountClient, err := account.NewClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		accountClient,
	}, nil
}

// 获取基本信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/Mini_Program_Information_Settings.html
func (comp *Client) GetBasicInfo(ctx context.Context) (*response.ResponseGetBasicInfo, error) {

	result := &response.ResponseGetBasicInfo{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/account/getaccountbasicinfo", nil, nil, nil, result)

	return result, err

}

// 修改头像
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifyheadimage.html
func (comp *Client) UpdateAvatar(
	ctx context.Context,
	mediaID string,
	left float32,
	top float32,
	right float32,
	bottom float32,
) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"head_img_media_id": mediaID,
		"x1":                fmt.Sprintf("%f", left),
		"y1":                fmt.Sprintf("%f", top),
		"x2":                fmt.Sprintf("%f", right),
		"y2":                fmt.Sprintf("%f", bottom),
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/account/modifyheadimage", params, nil, nil, result)

	return result, err

}

// 修改功能介绍
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modifysignature.html
func (comp *Client) UpdateSignature(ctx context.Context, signature string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"signature": signature,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/account/modifysignature", params, nil, nil, result)

	return result, err

}
