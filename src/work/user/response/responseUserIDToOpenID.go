package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	response.ResponseWork

	OpenID string `json:"openid"`
}
