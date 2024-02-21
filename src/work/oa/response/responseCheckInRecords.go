package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInRecords struct {
	response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
