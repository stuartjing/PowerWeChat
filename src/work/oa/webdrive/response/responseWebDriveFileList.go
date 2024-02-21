package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileList struct {
	response.ResponseWork

	HasMore   bool             `json:"has_more"`
	NextStart int              `json:"next_start"`
	FileList  []*power.HashMap `json:"file_list"`
}
