package notify

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/models"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/payment/notify/request"
	"net/http"
)

type Refund struct {
	*Handler
}

func NewRefundNotify(app kernel.ApplicationPaymentInterface, request *http.Request) *Refund {

	paid := &Refund{
		NewHandler(app, request),
	}

	return paid
}

func (comp *Refund) Handle(closure func(message *request.RequestNotify, refund *models.Refund, fail func(message string)) interface{}) (*http.Response, error) {

	message, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	reqInfo, err := comp.reqInfo()
	if err != nil {
		return nil, err
	}

	// struct the content
	refund := &models.Refund{}
	err = object.JsonDecode([]byte(reqInfo), refund)
	if err != nil {
		return nil, err
	}

	result := closure(message, refund, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()

}
