package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseApprovalCreate struct {
	response.ResponseWork
	SpNo string `json:"sp_no"`
}
