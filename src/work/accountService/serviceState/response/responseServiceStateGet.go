package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseServiceStateGet struct {
	response.ResponseWork

	ServiceState   int    `json:"service_state"`
	ServicerUserID string `json:"servicer_userid"`
}
