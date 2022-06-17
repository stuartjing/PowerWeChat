package jssdk

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/http/response"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Client struct {
	*kernel.BaseClient

	*kernel.InteractsWithCache

	TicketEndpoint string
	url            string
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient:         baseClient,
		InteractsWithCache: &kernel.InteractsWithCache{},
	}

	client.TicketEndpoint = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"

	return client, nil
}

func (comp *Client) BuildConfig(jsApiList []string, debug bool, beta bool, openTagList []string, url string) (interface{}, error) {

	signature, err := comp.ConfigSignature(url, "", 0)
	if err != nil {
		return signature, err
	}
	config := object.MergeHashMap(&object.HashMap{
		"debug":       debug,
		"beta":        beta,
		"jsApiList":   jsApiList,
		"openTagList": openTagList,
	}, signature)

	powerConfig, err := power.HashMapToPower(config)
	return powerConfig, err

}

func (comp *Client) GetConfigArray(apis []string, debug bool, beta bool, openTagList []string, url string) (string, error) {

	result, err := comp.BuildConfig(apis, debug, beta, openTagList, url)

	return result.(string), err
}

func (comp *Client) GetTicket(refresh bool, ticketType string) (*object.HashMap, error) {

	cacheKey := fmt.Sprintf("powerwechat.basic_service.jssdk.ticket.%s.%s", ticketType, comp.GetAppID())

	if !refresh && comp.GetCache().Has(cacheKey) {
		ticket, err := comp.GetCache().Get(cacheKey, nil)
		return ticket.(*object.HashMap), err
	}

	mapRSBody := &object.HashMap{}
	resultBody := ""
	rs, err := comp.RequestRaw(comp.TicketEndpoint, "GET", &object.HashMap{
		"query": &object.StringMap{
			"type": ticketType,
		}}, nil, &resultBody)
	if err != nil {
		return nil, err
	}

	err = object.JsonDecode([]byte(resultBody), mapRSBody)
	if (*mapRSBody)["errcode"].(float64) != 0 {
		return mapRSBody, errors.New((*mapRSBody)["errmsg"].(string))
	}

	result, err := comp.CastResponseToType(rs.(*response.HttpResponse).Response, response2.TYPE_MAP)
	if err != nil {
		return nil, err
	}

	resultData := result.(*object.HashMap)
	err = comp.GetCache().Set(cacheKey, result, time.Duration((*resultData)["expires_in"].(float64)-500)*time.Second)
	if err != nil {
		return nil, err
	}

	if !comp.GetCache().Has(cacheKey) {
		return nil, errors.New("Failed to cache jssdk ticket. ")
	}

	return resultData, nil
}

func (comp *Client) ConfigSignature(url string, nonce string, timestamp int64) (*object.HashMap, error) {

	if nonce == "" {
		nonce = object.QuickRandom(10)
	}
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}

	result, err := comp.GetTicket(false, "jsapi")
	if err != nil {
		return result, err
	}
	ticket := (*result)["ticket"].(string)

	return &object.HashMap{
		"appId":     comp.GetAppID(),
		"nonceStr":  nonce,
		"timestamp": timestamp,
		"url":       url,
		"signature": comp.GetTicketSignature(ticket, nonce, timestamp, url),
	}, nil

}

func (comp *Client) GetTicketSignature(ticket string, nonce string, timestamp int64, url string) string {

	param := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, timestamp, url)

	hash := sha1.New()
	hash.Write([]byte(param))
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (comp *Client) dictionaryOrderSignature(params []string) string {
	sort.Strings(params)
	strJoined := strings.Join(params, "")
	hash := sha1.New()
	hash.Write([]byte(strJoined))
	return string(hash.Sum(nil))
}

func (comp *Client) SetUrl(url string) *Client {

	comp.url = url

	return comp
}

func (comp *Client) GetUrl(externalRequest *http.Request) string {

	if comp.url != "" {
		return comp.url
	}
	return externalRequest.URL.String()
}

func (comp *Client) GetAppID() string {
	config := (*comp.App).GetConfig()
	return config.GetString("app_id", "")
}

func (comp *Client) getAgentID() string {
	config := (*comp.App).GetConfig()
	return config.GetString("agent_id", "")
}
