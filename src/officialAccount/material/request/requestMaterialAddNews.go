package request

import "github.com/stuartjing/PowerWeChat/v3/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}
