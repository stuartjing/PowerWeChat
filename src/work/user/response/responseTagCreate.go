package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagCreate struct {
	response.ResponseWork

	TagID string `json:"tagid"`
}
