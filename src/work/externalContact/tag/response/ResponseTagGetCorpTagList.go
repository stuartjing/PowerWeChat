package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagGetCorpTagList struct {
	response.ResponseWork

	TagGroups []*CorpTagGroup `json:"tag_group"`
}
