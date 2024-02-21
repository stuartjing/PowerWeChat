package promotion

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/stuartjing/PowerWeChat/v3/src/payment/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/promotion/request"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/promotion/response"
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

// 向员工付款
// https://work.weixin.qq.com/api/doc/90000/90135/90278
func (comp *Client) PayTransferToPocket(ctx context.Context, data *request.RequestPayTransferToPocket) (*response.ResponsePayTransferToPocket, error) {
	result := &response.ResponsePayTransferToPocket{}

	//params, err := object.StructToHashMapWithTag(data, "json")
	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/paywwsptrans2pocket")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}

// 查询付款记录
// https://work.weixin.qq.com/api/doc/90000/90135/90279
func (comp *Client) QueryTransferToPocket(ctx context.Context, data *request.RequestQueryTransferToPocket) (*response.ResponseQueryTransferToPocket, error) {

	result := &response.ResponseQueryTransferToPocket{}

	//params, err := object.StructToHashMapWithTag(data,"json")
	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/mmpaymkttransfers/promotion/querywwsptrans2pocket")
	_, err = comp.SafeRequest(ctx, endpoint, params, http.MethodPost, &object.HashMap{}, nil, result)

	return result, err
}
