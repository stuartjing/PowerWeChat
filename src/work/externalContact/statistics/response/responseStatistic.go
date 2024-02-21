package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseStatistic struct {
	response.ResponseWork

	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`

	Items []*power.HashMap `json:"items"`
}
