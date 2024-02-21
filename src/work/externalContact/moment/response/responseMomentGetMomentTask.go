package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentGetMomentTask struct {
	response.ResponseWork

	TaskList   []*power.HashMap `json:"task_list"`
	NextCursor string           `json:"next_cursor"`
}
