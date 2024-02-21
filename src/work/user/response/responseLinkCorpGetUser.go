package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`
}
