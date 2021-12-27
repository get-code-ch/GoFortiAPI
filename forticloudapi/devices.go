package forticloudapi

import (
	"encoding/json"
	"fmt"
	"github.com/get-code-ch/GoFortiAPI"
	"io"
	"io/ioutil"
)

type Devices struct {
	api      *GoFortiAPI.FortiAPI
	apiPath  string
	request  DevicesRequest
	response *DevicesList
}

type DevicesRequest []byte
type DevicesList []DeviceSummary

func (dr DevicesRequest) Body() []byte {
	return []byte{}
}

func (dl *DevicesList) Init(body io.ReadCloser) error {
	if b, err := ioutil.ReadAll(body); err != nil {
		return err
	} else {
		if err := json.Unmarshal(b, &dl); err != nil {
			return err
		}
	}
	return nil
}

func (dl *DevicesList) String() string {
	var devices string

	for _, d := range *dl {
		devices += fmt.Sprintf("%s\t%s\t%s\t%s\t%t\t%s\n\t%s\t%s\n", d.Sn, d.FirmwareVersion, d.Model, d.Ip, d.TunnelAlive, d.Name, d.Latitude, d.Longitude)
	}
	return devices
}

func NewDevicesList(api *GoFortiAPI.FortiAPI) *Devices {
	// Creating devicesGet
	devices := new(Devices)
	devices.api = api

	devices.apiPath = "devices"
	devices.request = []byte{}
	devices.response = new(DevicesList)

	return devices
}

func (d *Devices) Post() error {
	return d.api.PostRequest(d.apiPath, d.request, d.response)
}

func (d *Devices) Get() error {
	return d.api.GetRequest(d.apiPath, d.request, d.response)
}

func (d *Devices) List() *DevicesList {
	return d.response
}
