package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/models"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentGet struct {
	response.ResponseWork
	Department *models.Department `json:"department"`
}
