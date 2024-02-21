package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUserSimpleList struct {
	response.ResponseWork

	UserList []*power.HashMap `json:"userlist"`
}
