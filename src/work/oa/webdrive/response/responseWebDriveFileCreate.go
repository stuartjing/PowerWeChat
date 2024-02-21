package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileCreate struct {
	response.ResponseWork

	FileID string `json:"fileid"`
	Url    string `json:"url"`
}
