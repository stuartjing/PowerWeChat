package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	response.ResponseWork
	ID int `json:"id"`
}
