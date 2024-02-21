package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileMove struct {
	response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}
