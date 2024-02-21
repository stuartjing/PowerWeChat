package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagGetStrategyTagList struct {
	response.ResponseWork

	TagGroups []*StrategyTagGroup `json:"tag_group"`
}
