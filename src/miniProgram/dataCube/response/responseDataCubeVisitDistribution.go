package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeVisit struct {
	response.ResponseMiniProgram

	RefDate string           `json:"ref_date"`
	List    []*power.HashMap `json:"list"`
}
