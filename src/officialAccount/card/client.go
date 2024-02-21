package card

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	response2 "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
	"github.com/stuartjing/PowerWeChat/v3/src/officialAccount/card/request"
	"github.com/stuartjing/PowerWeChat/v3/src/officialAccount/card/response"
	"strings"
)

type Client struct {
	*kernel.BaseClient

	URL               string
	TicketCacheKey    string
	TicketCachePrefix string
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
	}

	client.TicketCachePrefix = "powerwechat.official_account.card.api_ticket."

	return client, nil

}

func (comp *Client) Colors() {

}

func (comp *Client) Categories() {

}

// 创建卡券
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html
func (comp *Client) Create(ctx context.Context, param *request.RequestCardCreate) (*response.ResponseCardCreate, error) {
	result := &response.ResponseCardCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/create", param, nil, nil, result)

	return result, err

}

// 查看卡券详情
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) Get(ctx context.Context, cardID string) (*response.ResponseCardGet, error) {
	result := &response.ResponseCardGet{}

	param := object.HashMap{
		"card_id": cardID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/get", param, nil, nil, result)

	return result, err

}

// 批量查询卡券列表
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) List(ctx context.Context, offset int, count int, statusList string) (*response.ResponseCardList, error) {
	result := &response.ResponseCardList{}

	param := object.HashMap{
		"offset":      offset,
		"count":       count,
		"status_list": statusList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/batchget", param, nil, nil, result)

	return result, err

}

// 更改卡券信息接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) Update(ctx context.Context, cardID string, cardType string, card interface{}) (*response.ResponseCardUpdate, error) {
	result := &response.ResponseCardUpdate{}

	param := object.HashMap{
		"card_id":                 cardID,
		strings.ToLower(cardType): card,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/update", param, nil, nil, result)

	return result, err

}

// 删除卡券接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) Delete(ctx context.Context, cardID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	param := object.HashMap{
		"card_id": cardID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/delete", param, nil, nil, result)

	return result, err

}

// 创建二维码接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html#0
func (comp *Client) CreateQrCode(ctx context.Context, param request.RequestCreateQrCode) (*response.ResponseCreateQrCode, error) {
	result := &response.ResponseCreateQrCode{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/qrcode/create", param, nil, nil, result)

	return result, err

}
