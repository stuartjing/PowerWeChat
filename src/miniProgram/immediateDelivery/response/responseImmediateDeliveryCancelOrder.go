package response

import "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"

type ResponseImmediateDeliveryCancelOrder struct {
	response.ResponseMiniProgram

	DeductFee int64  `json:"deduct_fee"`
	Desc      string `json:"desc"`
}
