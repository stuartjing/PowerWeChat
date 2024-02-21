package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
