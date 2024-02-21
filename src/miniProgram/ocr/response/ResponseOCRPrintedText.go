package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseOCRPrintedText struct {
	response.ResponseMiniProgram
	Items   []*power.HashMap `json:"items"`
	ImgSize []*power.HashMap `json:"img_size"`
}
