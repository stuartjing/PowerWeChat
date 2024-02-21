package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
