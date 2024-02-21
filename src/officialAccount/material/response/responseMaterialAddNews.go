package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMaterialAddNews struct {
	response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
}
