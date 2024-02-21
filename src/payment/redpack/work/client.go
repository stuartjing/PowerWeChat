package work

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/stuartjing/PowerWeChat/v3/src/payment/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/redpack/request"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/redpack/response"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// Send Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90275
func (comp *Client) SendWorkWX(ctx context.Context, data *request.RequestSendWorkWX) (*response.ResponseSendWorkWX, error) {
	result := &response.ResponseSendWorkWX{}
	if data.WXAppID == "" {
		config := (*comp.App).GetConfig()
		data.WXAppID = config.GetString("app_id", "")
	}

	//params, err := object.StructToHashMapWithTag(data,"json")
	params, err := object.StructToHashMap(data)

	endpoint := comp.Wrap("/mmpaymkttransfers/sendworkwxredpack")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// Query Work WX redpack.
// https://work.weixin.qq.com/api/doc/90000/90135/90276
func (comp *Client) QueryWorkWX(ctx context.Context, data *request.RequestQueryWorkWX) (*response.ResponseQueryWorkWX, error) {
	result := &response.ResponseQueryWorkWX{}

	if data.Appid == "" {
		config := (*comp.App).GetConfig()
		data.Appid = config.GetString("app_id", "")
	}

	//params, err := object.StructToHashMapWithTag(data,"json")
	params, err := object.StructToHashMap(data)

	endpoint := comp.Wrap("/mmpaymkttransfers/queryworkwxredpack")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}
