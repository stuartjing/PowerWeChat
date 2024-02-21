package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpGroupGetToken struct {
	response.ResponseWork
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
