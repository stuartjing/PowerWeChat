package request

import "github.com/stuartjing/PowerWeChat/v3/src/kernel/power"

type RequestCheckInSetScheduleList struct {
	GroupID   int              `json:"groupid"`
	Items     []*power.HashMap `json:"items"`
	YearMonth int              `json:"yearmonth"`
}
