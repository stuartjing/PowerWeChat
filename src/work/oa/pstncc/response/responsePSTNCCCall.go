package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponsePSTNCCCall struct {
	response.ResponseWork

	States []*power.HashMap `json:"states"`
}
