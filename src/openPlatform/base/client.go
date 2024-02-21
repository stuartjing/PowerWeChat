package base

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	response2 "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/base/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 使用授权码获取授权信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html
func (comp *Client) HandleAuthorize(ctx context.Context, authCode string) (*response.ResponseHandleAuthorize, error) {

	result := &response.ResponseHandleAuthorize{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.HashMap{
		"component_appid":    config.GetString("app_id", ""),
		"authorization_code": authCode,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/component/api_query_auth", params, nil, nil, result)

	return result, err
}

// 获取授权帐号详情
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_get_authorizer_info.html
func (comp *Client) GetAuthorizer(ctx context.Context, appID string) (*response.ResponseGetAuthorizer, error) {

	result := &response.ResponseGetAuthorizer{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.HashMap{
		"component_appid":  config.GetString("app_id", ""),
		"authorizer_appid": appID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/component/api_get_authorizer_info", params, nil, nil, result)

	return result, err
}

// 获取授权方选项信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/Account_Authorization/api_get_authorizer_option.html
func (comp *Client) GetAuthorizerOption(ctx context.Context, appID string, name string) (*response.ResponseGetAuthorizerOption, error) {

	result := &response.ResponseGetAuthorizerOption{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.HashMap{
		"component_appid":  config.GetString("app_id", ""),
		"authorizer_appid": appID,
		"option_name":      name,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/component/api_get_authorizer_option", params, nil, nil, result)

	return result, err
}

// 设置授权方选项信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/Account_Authorization/api_set_authorizer_option.html
func (comp *Client) SetAuthorizerOption(ctx context.Context, appID string, name string, value string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.HashMap{
		"component_appid":  config.GetString("app_id", ""),
		"authorizer_appid": appID,
		"option_name":      name,
		"option_value":     value,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/component/api_set_authorizer_option", params, nil, nil, result)

	return result, err
}

// 拉取所有已授权的帐号信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/Account_Authorization/api_get_authorizer_list.html
func (comp *Client) GetAuthorizers(ctx context.Context, offset int, count int) (*response.ResponseGetAuthorizers, error) {

	result := &response.ResponseGetAuthorizers{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.HashMap{
		"component_appid": config.GetString("app_id", ""),
		"offset":          offset,
		"count":           count,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/component/api_get_authorizer_list", params, nil, nil, result)

	return result, err
}

// 预授权码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
func (comp *Client) CreatePreAuthorizationCode(ctx context.Context) (*response.ResponseCreatePreAuthorizationCode, error) {

	result := &response.ResponseCreatePreAuthorizationCode{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.HashMap{
		"component_appid": config.GetString("app_id", ""),
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/component/api_create_preauthcode", params, nil, nil, result)

	return result, err
}

// 清空 api 的调用quota
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/openApi/clear_quota.html
func (comp *Client) ClearQuota(ctx context.Context) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	config := (*comp.BaseClient.App).GetConfig()

	params := &object.HashMap{
		"component_appid": config.GetString("app_id", ""),
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/component/clear_quota", params, nil, nil, result)

	return result, err
}
