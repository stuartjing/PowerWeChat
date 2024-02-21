package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	response.ResponseWork

	UserID string `json:"userid"`
}
