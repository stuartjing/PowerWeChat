package response

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMessageSyncMsg struct {
	response.ResponseWork

	NextCursor string            `json:"next_cursor"`
	HasMore    int               `json:"has_more"`
	MsgList    []*object.HashMap `json:"msg_list"`
}
