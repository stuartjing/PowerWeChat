package miniProgram

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/providers"
	"github.com/stuartjing/PowerWeChat/v3/src/miniProgram"
	"github.com/stuartjing/PowerWeChat/v3/src/officialAccount/material"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/account"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/auth"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/code"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/domain"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/privacy"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/setting"
	"github.com/stuartjing/PowerWeChat/v3/src/openPlatform/authorizer/miniProgram/tester"
)

type Application struct {
	*miniProgram.MiniProgram

	Auth     *auth.Client
	Account  *account.Client
	Code     *code.Client
	Domain   *domain.Client
	Material *material.Client
	Privacy  *privacy.Client
	Setting  *setting.Client
	Tester   *tester.Client
}

func NewApplication(config *miniProgram.UserConfig, extraInfos ...*kernel.ExtraInfo) (*Application, error) {

	var extraInfo, _ = kernel.NewExtraInfo()
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	miniProgram, err := miniProgram.NewMiniProgram(config, extraInfo)
	if err != nil {
		return nil, err
	}
	app := &Application{
		MiniProgram: miniProgram,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Aggregate --------------
	app.Account, err = account.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Code --------------
	app.Code, err = code.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Domain --------------
	app.Domain, err = domain.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Material --------------
	app.Material, err = material.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Privacy --------------
	app.Privacy, err = privacy.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Setting --------------
	app.Setting, err = setting.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Tester --------------
	app.Tester, err = tester.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	return app, err
}
