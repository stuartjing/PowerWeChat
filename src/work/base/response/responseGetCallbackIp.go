package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetCallBackIP struct {
	IPList []string `json:"ip_list"`
	response.ResponseWork
}
