package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseShortLinkGenerate struct {
	response.ResponseMiniProgram

	Link string `json:"link"`
}
