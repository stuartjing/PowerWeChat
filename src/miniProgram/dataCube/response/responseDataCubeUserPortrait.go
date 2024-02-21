package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeUserPortrait struct {
	response.ResponseMiniProgram

	RefDate    string         `json:"ref_date"`
	VisitUVNew *power.HashMap `json:"visit_uv_new"`
	VisitUV    *power.HashMap `json:"visit_uv"`
}
