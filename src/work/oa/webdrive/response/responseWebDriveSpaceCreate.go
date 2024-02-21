package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveSpaceCreate struct {
	response.ResponseWork

	SpaceID string `json:"spaceid"`
}
