package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileUpload struct {
	response.ResponseWork

	FileID string `json:"fileid"`
}
