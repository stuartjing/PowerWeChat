package auth

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/auth"
	"github.com/stuartjing/PowerWeChat/v3/src/work/miniProgram/response"
)

type Client struct {
	*kernel.BaseClient

	// PowerWechat\OpenPlatform\Application
	component kernel.ApplicationInterface
}

func NewClient(app kernel.ApplicationInterface, component kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
		component,
	}, nil
}

// 小程序登录
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/others/WeChat_login.html
func (comp *Client) Session(ctx context.Context, code string) (*response.ResponseSession, error) {

	result := &response.ResponseSession{}

	config := (*comp.App).GetConfig()
	componentConfig := comp.component.GetConfig()
	token := comp.component.GetComponent("AccessToken").(*auth.AccessToken)
	componentToken, err := token.GetToken(false)

	query := &object.StringMap{
		"appid":                  config.GetString("app_id", ""),
		"js_code":                code,
		"grant_type":             "authorization_code",
		"component_appid":        componentConfig.GetString("app_id", ""),
		"component_access_token": componentToken.ComponentAccessToken,
	}
	_, err = comp.BaseClient.HttpGet(ctx, "sns/component/jscode2session", query, nil, result)

	return result, err

}
