package response

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/models"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml

type ResponseOrder struct {
	response.ResponsePayment

	Amount          *models.TransactionAmount    `json:"amount"`
	AppID           string                       `json:"appid"`
	Attach          string                       `json:"attach"`
	BankType        string                       `json:"bank_type"`
	MchID           string                       `json:"mchid"`
	OutTradeNo      string                       `json:"out_trade_no"`
	Payer           *models.TransactionPayer     `json:"payer"`
	PromotionDetail []*object.HashMap            `json:"promotion_detail"`
	SuccessTime     string                       `json:"success_time"`
	TradeState      string                       `json:"trade_state"`
	TradeStateDesc  string                       `json:"trade_state_desc"`
	TradeType       string                       `json:"trade_type"`
	TransactionID   string                       `json:"transaction_id"`
	SceneInfo       *models.TransactionSceneInfo `json:"scene_info"`
}
