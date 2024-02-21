package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentGetMomentCustomerList struct {
	response.ResponseWork

	CustomerList []*power.HashMap `json:"customer_list"`
	NextCursor   string           `json:"next_cursor"`
}
