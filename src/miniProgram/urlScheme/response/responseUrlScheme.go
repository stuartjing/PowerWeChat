package response

import "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"

type ResponseUrlScheme struct {
	response.ResponseMiniProgram

	OpenLink string `json:"openlink"`
}
