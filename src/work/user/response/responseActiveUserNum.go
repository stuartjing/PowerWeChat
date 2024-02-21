package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserActiveCount struct {
	response.ResponseWork

	ActiveCount string `json:"active_cnt"`
}
