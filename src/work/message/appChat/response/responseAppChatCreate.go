package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseAppChatCreate struct {
	response.ResponseWork

	ChatID string `json:"chatid"`
}
