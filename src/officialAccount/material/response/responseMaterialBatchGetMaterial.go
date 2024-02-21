package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMaterialBatchGetMaterial struct {
	response.ResponseOfficialAccount

	TotalCount int              `json:"total_count"`
	ItemCount  int              `json:"item_count"`
	Item       []*power.HashMap `json:"item"`
}
