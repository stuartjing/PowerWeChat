package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInSchedulist struct {
	response.ResponseWork

	ScheduleList []*power.HashMap `json:"schedule_list"`
}
