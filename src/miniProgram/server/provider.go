package server

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/messages"
	"github.com/stuartjing/PowerWeChat/v3/src/officialAccount/server"
	"github.com/stuartjing/PowerWeChat/v3/src/work/server/handlers"
)

func RegisterProvider(app kernel.ApplicationInterface) (*server.Guard, error) {

	guard := server.NewGuard(&app)
	echoStrHandler := handlers.NewEchoStrHandler(&app)
	guard.Push(echoStrHandler, messages.VOID)

	return guard, nil
}
