package forticloudapi

import (
	"encoding/json"
	"fmt"
	"github.com/get-code-ch/GoFortiAPI"
	"io"
	"io/ioutil"
)

type Device struct {
	api      *GoFortiAPI.FortiAPI
	apiPath  string
	request  DeviceRequest
	response *DeviceSummary
}

type DeviceRequest []byte

type DeviceSummary struct {
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

func (dr DeviceRequest) Body() []byte {
	return []byte{}
}

func (d *DeviceSummary) Init(body io.ReadCloser) error {
	if b, err := ioutil.ReadAll(body); err != nil {
		return err
	} else {
		if err := json.Unmarshal(b, &d); err != nil {
			return err
		}
	}
	return nil
}

func (d *DeviceSummary) String() string {

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%t\t%s\n\t%s\t%s\n",
		d.Sn, d.FirmwareVersion, d.Model, d.Ip, d.TunnelAlive, d.Name, d.Latitude, d.Longitude)
}

func NewDevice(api *GoFortiAPI.FortiAPI, serial string) *Device {
	// Creating devicesGet
	device := new(Device)
	device.api = api

	device.apiPath = "devices/" + serial
	device.request = []byte{}
	device.response = new(DeviceSummary)

	return device
}

func (d *Device) Post() error {
	return d.api.PostRequest(d.apiPath, d.request, d.response)
}

func (d *Device) Get() error {
	return d.api.GetRequest(d.apiPath, d.request, d.response)
}

func (d *Device) Item() *DeviceSummary {
	return d.response
}
