package response

import (
	"encoding/xml"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseReserve struct {
	response.ResponsePayment

	XMLName  xml.Name `xml:"xml"`
	Text     string   `xml:",chardata"`
	AppID    string   `xml:"appid"`
	MchID    string   `xml:"mch_id"`
	NonceStr string   `xml:"nonce_str"`
	Sign     string   `xml:"sign"`
}
