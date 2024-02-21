package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileRename struct {
	response.ResponseWork

	File *power.HashMap `json:"file"`
}
