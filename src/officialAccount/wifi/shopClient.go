package wifi

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel"
	response2 "github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
	"github.com/stuartjing/PowerWeChat/v3/src/officialAccount/wifi/response"
)

type ShopClient struct {
	*kernel.BaseClient
}

func NewShopClient(app kernel.ApplicationInterface) (*ShopClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &ShopClient{
		baseClient,
	}, nil
}

// 查询门店 WiFi 信息接口
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/WiFi_mini_programs.html
func (comp *ShopClient) Get(ctx context.Context, shopID int) (*response.ResponseWifiShopGet, error) {

	result := &response.ResponseWifiShopGet{}

	params := &object.HashMap{
		"shop_id": shopID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/shop/get", params, nil, nil, result)

	return result, err

}

// 通过此接口获取 WiFi 的门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Stores_management/Get_a_Wi-Fi_store_list.html
func (comp *ShopClient) List(ctx context.Context, page int, size int) (*response.ResponseWifiShopList, error) {

	result := &response.ResponseWifiShopList{}

	params := &object.HashMap{
		"pageindex": page,
		"pagesize":  size,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/shop/List", params, nil, nil, result)

	return result, err

}

// 通过此接口修改门店的网络信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Stores_management/Modify_the_store_network_information.html
func (comp *ShopClient) Update(ctx context.Context, shopID int, oldSSID string, ssid string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"shop_id":  shopID,
		"old_ssid": oldSSID,
		"ssid":     ssid,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/shop/update", params, nil, nil, result)

	return result, err
}

// 通过此接口清空门店的网络配置及所有设备，恢复空门店状态。
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Stores_management/Empty_the_store_network_and_devices.html
func (comp *ShopClient) ClearDevice(ctx context.Context, shopID int, ssid string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"shop_id": shopID,
	}

	if ssid != "" {
		(*params)["ssid"] = ssid
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/shop/clean", params, nil, nil, result)

	return result, err
}
