package response

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseImmediateDeliveryDelivery struct {
	response.ResponseMiniProgram

	List []*object.HashMap `json:"list"`
}
