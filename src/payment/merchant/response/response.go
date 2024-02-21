package response

import "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"

type ResponseMediaUpload struct {
	response.ResponsePayment
	MediaId string `json:"media_id"`
}
