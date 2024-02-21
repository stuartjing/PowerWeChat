package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseJoinCode struct {
	response.ResponseWork

	JoinCode string `json:"join_qrcode"`
}
