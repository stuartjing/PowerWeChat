package codeTemplate

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	response2 "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/codeTemplate/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取代码草稿列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html
func (comp *Client) GetDrafts(ctx context.Context) (*response.ResponseGetDrafts, error) {

	result := &response.ResponseGetDrafts{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/gettemplatedraftlist", nil, nil, result)

	return result, err

}

// 将草稿添加到代码模板库
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html#请求地址
func (comp *Client) CreateFromDraft(ctx context.Context, draftID int, templateType int) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"draft_id":      draftID,
		"template_type": templateType,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/addtotemplate", params, nil, nil, result)

	return result, err

}

// 获取代码模板列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatelist.html#请求地址
func (comp *Client) List(ctx context.Context) (*response.ResponseList, error) {

	result := &response.ResponseList{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/gettemplatelist", nil, nil, result)

	return result, err

}

// 删除指定代码模板
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html#请求地址
func (comp *Client) Delete(ctx context.Context, templateID string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"template_id": templateID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/deletetemplate", params, nil, nil, result)

	return result, err

}
