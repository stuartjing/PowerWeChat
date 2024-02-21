package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeGetPerformanceData struct {
	Data *power.HashMap `json:"data"`

	response.ResponseMiniProgram
}
