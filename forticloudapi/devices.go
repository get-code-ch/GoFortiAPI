package forticloudapi

import (
	"encoding/json"
	"fmt"
	"github.com/get-code-ch/GoFortiAPI"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//TODO: Define body to POST
/*
type DevicesPostBody struct {
}
type DevicesPost struct {
	api      *FortiAPI
	url      url.URL
	header   http.Header
	body     []byte
	Response *DevicesPostResponse
}
*/

type DevicesGet struct {
	api      *GoFortiAPI.FortiAPI
	url      url.URL
	header   http.Header
	body     []byte
	Response *DevicesGetResponse
}

type DevicesGetResponse []struct {
	Sn              string      `json:"sn"`
	Name            string      `json:"name"`
	TimeZone        json.Number `json:"timeZone"`
	TunnelAlive     bool        `json:"tunnelAlive"`
	ContractEndTime int         `json:"contractEndTime"`
	Model           string      `json:"model"`
	FirmwareVersion string      `json:"firmwareVersion"`
	Management      bool        `json:"management"`
	Initialized     bool        `json:"initialized"`
	SubAccountOid   int         `json:"subAccountOid"`
	Ip              string      `json:"ip"`
	Latitude        string      `json:"latitude"`
	Longitude       string      `json:"longitude"`
	Total           int         `json:"total"`
	Trial           bool        `json:"trial"`
}

func NewDevicesGetResponse() *DevicesGetResponse {
	return new(DevicesGetResponse)
}

func (r *DevicesGetResponse) Init(body io.ReadCloser) error {
	if b, err := ioutil.ReadAll(body); err != nil {
		return err
	} else {
		if b[0] == '{' {
			b = append([]byte{'['}, b...)
			b = append(b, ']')
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return err
		}
	}
	return nil
}

func (r *DevicesGetResponse) String() string {
	var devices string

	for _, d := range *r {
		devices += fmt.Sprintf("%s\t%s\t%s\t%s\t%t\t%s\n\t%s\t%s\n", d.Sn, d.FirmwareVersion, d.Model, d.Ip, d.TunnelAlive, d.Name, d.Latitude, d.Longitude)
	}
	return devices
}

func NewDevicesGet(api *GoFortiAPI.FortiAPI, token GoFortiAPI.AccessToken) *DevicesGet {
	// Creating devicesGet
	devicesGet := new(DevicesGet)

	// Setting URL to API
	devicesGet.url.Scheme = "https"
	devicesGet.url.Host = api.Region
	devicesGet.url.Path = "forticloudapi/v1/devices"

	// Setting Header
	devicesGet.header = http.Header{}
	devicesGet.header.Set("Content-Type", "application/json")
	devicesGet.header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	devicesGet.Response = NewDevicesGetResponse()

	// Setting AuthPost Body
	devicesGet.body = nil
	devicesGet.api = api

	return devicesGet
}

func (request *DevicesGet) API() GoFortiAPI.APICall {
	return GoFortiAPI.APICall{Header: request.header, URL: request.url.String(), Body: request.body}
}

func (request *DevicesGet) Post() error {
	return request.api.PostRequest(request, request.Response)
}

func (request *DevicesGet) Get() error {
	return request.api.GetRequest(request, request.Response)
}

func (request *DevicesGet) GetSn(sn string) error {
	basePath := request.url.Path
	request.url.Path = basePath + "/" + sn
	err := request.api.GetRequest(request, request.Response)
	request.url.Path = basePath
	return err
}
