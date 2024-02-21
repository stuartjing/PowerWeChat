package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetLivingCode struct {
	response.ResponseWork

	LivingCode int `json:"living_code"`
}
