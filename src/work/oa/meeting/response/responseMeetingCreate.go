package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingCreate struct {
	response.ResponseWork

	MeetingID int `json:"meetingid"`
}
