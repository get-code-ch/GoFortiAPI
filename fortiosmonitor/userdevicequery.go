package fortiosmonitor

import (
	"encoding/json"
	"fmt"
	"github.com/get-code-ch/GoFortiAPI"
	"io"
	"io/ioutil"
)

type UserDeviceQuery struct {
	api      *GoFortiAPI.FortiAPI
	apiPath  string
	request  UserDeviceQueryRequest
	response *UserDeviceQueryResponse
}

type UserDeviceQueryRequest []byte

type UserDeviceQueryResponse struct {
	HttpMethod string `json:"http_method"`
	Results    []struct {
		Ipv4Address                     string   `json:"ipv4_address"`
		Mac                             string   `json:"mac"`
		HardwareVendor                  string   `json:"hardware_vendor,omitempty"`
		HardwareVersion                 string   `json:"hardware_version,omitempty"`
		HardwareType                    string   `json:"hardware_type,omitempty"`
		HardwareFamily                  string   `json:"hardware_family,omitempty"`
		Vdom                            string   `json:"vdom"`
		OsName                          string   `json:"os_name"`
		OsVersion                       string   `json:"os_version,omitempty"`
		Hostname                        string   `json:"hostname"`
		LastSeen                        int      `json:"last_seen"`
		HostSrc                         string   `json:"host_src"`
		UnjoinedForticlientEndpoint     bool     `json:"unjoined_forticlient_endpoint"`
		IsOnline                        bool     `json:"is_online"`
		ActiveStartTime                 int      `json:"active_start_time"`
		IsFortiguardSrc                 bool     `json:"is_fortiguard_src"`
		MasterMac                       string   `json:"master_mac"`
		DetectedInterface               string   `json:"detected_interface"`
		MacFirewallAddress              string   `json:"mac_firewall_address,omitempty"`
		IsMasterDevice                  bool     `json:"is_master_device"`
		IsDetectedInterfaceRoleWan      bool     `json:"is_detected_interface_role_wan"`
		DetectedInterfaceFortitelemetry bool     `json:"detected_interface_fortitelemetry"`
		DhcpLeaseStatus                 string   `json:"dhcp_lease_status,omitempty"`
		DhcpLeaseExpire                 int      `json:"dhcp_lease_expire,omitempty"`
		DhcpLeaseReserved               bool     `json:"dhcp_lease_reserved,omitempty"`
		DhcpServerId                    int      `json:"dhcp_server_id,omitempty"`
		OnlineInterfaces                []string `json:"online_interfaces,omitempty"`
	} `json:"results"`
	Vdom      string `json:"vdom"`
	QueryType string `json:"query_type"`
	Count     int    `json:"count"`
	Total     int    `json:"total"`
	Start     int    `json:"start"`
	Number    int    `json:"number"`
	Status    string `json:"status"`
	Serial    string `json:"serial"`
	Version   string `json:"version"`
	Build     int    `json:"build"`
}

func (udqReq UserDeviceQueryRequest) Body() []byte {
	return []byte{}
}

func (udqRsp *UserDeviceQueryResponse) Init(body io.ReadCloser) error {
	if b, err := ioutil.ReadAll(body); err != nil {
		return err
	} else {
		if err := json.Unmarshal(b, &udqRsp); err != nil {
			return err
		}
	}
	return nil
}

func (udqRsp *UserDeviceQueryResponse) String() string {
	var users string

	for _, u := range udqRsp.Results {
		users += fmt.Sprintf("%s\t%s\t%s\n", u.Hostname, u.HardwareType, u.DetectedInterface)
	}
	return users
}

func NewUserDeviceQuery(api *GoFortiAPI.FortiAPI, serial string) *UserDeviceQuery {
	// Creating userDeviceQuery
	userDeviceQuery := new(UserDeviceQuery)
	userDeviceQuery.api = api

	// Setting URL to API
	userDeviceQuery.apiPath = "fgt/" + serial + "/api/v2/monitor/user/device/query"
	userDeviceQuery.request = []byte{}
	userDeviceQuery.response = new(UserDeviceQueryResponse)

	return userDeviceQuery
}

func (udq *UserDeviceQuery) Post() error {
	return udq.api.PostRequest(udq.apiPath, udq.request, udq.response)
}

func (udq *UserDeviceQuery) Get() error {
	return udq.api.GetRequest(udq.apiPath, udq.request, udq.response)
}
