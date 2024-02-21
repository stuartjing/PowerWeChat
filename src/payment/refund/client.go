package refund

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	payment "github.com/stuartjing/PowerWeChat/v3/src/payment/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/refund/request"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/refund/response"
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

// Refund.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml

func (comp *Client) Refund(ctx context.Context, options *request.RequestRefund) (*response.ResponseRefund, error) {

	result := &response.ResponseRefund{}

	//body, err := object.StructToHashMapWithTag(options, "json")
	body, err := object.StructToHashMap(options)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/v3/refund/domestic/refunds")
	_, err = comp.PlainRequest(ctx, endpoint, nil, http.MethodPost, body, false, nil, result)

	return result, err
}

// Query Refund.
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_10.shtml

func (comp *Client) Query(ctx context.Context, refundOutNO string) (*response.ResponseRefund, error) {

	result := &response.ResponseRefund{}

	endpoint := fmt.Sprintf("/v3/refund/domestic/refunds/%s", refundOutNO)
	endpoint = comp.Wrap(endpoint)
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}
