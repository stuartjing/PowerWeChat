package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseScheduleAdd struct {
	response.ResponseWork

	ScheduleID string `json:"schedule_id"`
}
