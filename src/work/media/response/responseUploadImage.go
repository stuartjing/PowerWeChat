package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseUploadImage struct {
	response.ResponseWork

	URL string `json:"url"`
}
