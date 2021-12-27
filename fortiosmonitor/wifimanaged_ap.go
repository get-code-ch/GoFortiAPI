package fortiosmonitor

import (
	"encoding/json"
	"fmt"
	"github.com/get-code-ch/GoFortiAPI"
	"io"
	"io/ioutil"
)

type WifiManagedAP struct {
	api      *GoFortiAPI.FortiAPI
	apiPath  string
	request  WifiManagedAPRequest
	response *WifiManagedAPResponse
}

type WifiManagedAPRequest []byte

type WifiManagedAPResponse struct {
	HttpMethod string `json:"http_method"`
	Results    []struct {
		Name                  string `json:"name"`
		IsLocal               bool   `json:"is_local"`
		Vdom                  string `json:"vdom"`
		Serial                string `json:"serial"`
		ApProfile             string `json:"ap_profile"`
		State                 string `json:"state"`
		ConnectingFrom        string `json:"connecting_from"`
		ConnectingInterface   string `json:"connecting_interface"`
		Status                string `json:"status"`
		WtpId                 string `json:"wtp_id"`
		RegionCode            string `json:"region_code"`
		MgmtVlanid            int    `json:"mgmt_vlanid"`
		MeshUplink            string `json:"mesh_uplink"`
		MeshHopCount          int    `json:"mesh_hop_count"`
		MeshUplinkIntf        string `json:"mesh_uplink_intf"`
		MeshUplinkIntfSpeed   int    `json:"mesh_uplink_intf_speed"`
		Clients               int    `json:"clients"`
		OsVersion             string `json:"os_version"`
		LocalIpv4Addr         string `json:"local_ipv4_addr"`
		BoardMac              string `json:"board_mac"`
		JoinTime              string `json:"join_time"`
		JoinTimeRaw           int    `json:"join_time_raw"`
		ConnectionState       string `json:"connection_state"`
		ImageDownloadProgress int    `json:"image_download_progress"`
		LastFailure           string `json:"last_failure"`
		LastFailureCode       int    `json:"last_failure_code"`
		LastFailureParam      string `json:"last_failure_param"`
		LastFailureTime       string `json:"last_failure_time"`
		OverrideProfile       bool   `json:"override_profile"`
		Ssid                  []struct {
			Radio int           `json:"radio"`
			List  []interface{} `json:"list"`
		} `json:"ssid"`
		DataChanSec          string `json:"data_chan_sec"`
		DedicatedScanEnabled bool   `json:"dedicated_scan_enabled"`
		IndoorOutdoor        int    `json:"indoor_outdoor"`
		Subtype              int    `json:"subtype"`
		SensorsTemperatures  []int  `json:"sensors_temperatures"`
		Radio                []struct {
			RadioId                     int      `json:"radio_id"`
			Mode                        string   `json:"mode"`
			SupportedBands              []string `json:"supported_bands"`
			RadioType                   string   `json:"radio_type"`
			CountryName                 string   `json:"country_name,omitempty"`
			CountryCode                 int      `json:"country_code,omitempty"`
			ClientCount                 int      `json:"client_count,omitempty"`
			BaseBssid                   string   `json:"base_bssid,omitempty"`
			MaxVaps                     int      `json:"max_vaps,omitempty"`
			OperChan                    int      `json:"oper_chan,omitempty"`
			OperTxpower                 int      `json:"oper_txpower,omitempty"`
			ChannelUtilizationTimestamp int      `json:"channel_utilization_timestamp,omitempty"`
			ChannelUtilizationPercent   int      `json:"channel_utilization_percent,omitempty"`
			NoiseFloor                  int      `json:"noise_floor,omitempty"`
			BandwidthRx                 int      `json:"bandwidth_rx,omitempty"`
			BandwidthTx                 int      `json:"bandwidth_tx,omitempty"`
			BytesRx                     int      `json:"bytes_rx,omitempty"`
			BytesTx                     int      `json:"bytes_tx,omitempty"`
			InterferingAps              int      `json:"interfering_aps,omitempty"`
			TxRetriesPercent            int      `json:"tx_retries_percent,omitempty"`
			MacErrorsRx                 int      `json:"mac_errors_rx,omitempty"`
			MacErrorsTx                 int      `json:"mac_errors_tx,omitempty"`
			BackgroundScanEnabled       bool     `json:"background_scan_enabled,omitempty"`
			DetectInterfering           bool     `json:"detect_interfering,omitempty"`
			StaLocate                   bool     `json:"sta_locate,omitempty"`
			ChannelUtilization          bool     `json:"channel_utilization,omitempty"`
			OverrideBand                bool     `json:"override_band,omitempty"`
			OverrideTxpower             bool     `json:"override_txpower,omitempty"`
			AutoTxpower                 bool     `json:"auto_txpower,omitempty"`
			Txpower                     int      `json:"txpower,omitempty"`
			OverrideSa                  bool     `json:"override_sa,omitempty"`
			OverrideVaps                bool     `json:"override_vaps,omitempty"`
			Ssid                        struct {
			} `json:"ssid,omitempty"`
			AllSsids        bool          `json:"all_ssids,omitempty"`
			SsidMode        string        `json:"ssid_mode,omitempty"`
			OverrideChannel bool          `json:"override_channel,omitempty"`
			Channels        []interface{} `json:"channels,omitempty"`
			Health          struct {
				ChannelUtilization struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"channel_utilization"`
				ClientCount struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"client_count"`
				InterferingSsids struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"interfering_ssids"`
				InfraInterferingSsids struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"infra_interfering_ssids"`
				Overall struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"overall"`
			} `json:"health,omitempty"`
			Min2GChannel            int `json:"min_2g_channel,omitempty"`
			Min2GChannelUtilization int `json:"min_2g_channel_utilization,omitempty"`
			Max2GChannel            int `json:"max_2g_channel,omitempty"`
			Max2GChannelUtilization int `json:"max_2g_channel_utilization,omitempty"`
			Min5GChannel            int `json:"min_5g_channel,omitempty"`
			Min5GChannelUtilization int `json:"min_5g_channel_utilization,omitempty"`
			Max5GChannel            int `json:"max_5g_channel,omitempty"`
			Max5GChannelUtilization int `json:"max_5g_channel_utilization,omitempty"`
		} `json:"radio"`
		Health struct {
			General struct {
				CountryCode struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"country_code"`
				UplinkStatus []struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"uplink_status"`
				Overall struct {
					Value    int    `json:"value"`
					Severity string `json:"severity"`
				} `json:"overall"`
			} `json:"general"`
		} `json:"health"`
		Wired []struct {
			Interface  string `json:"interface"`
			BytesRx    int    `json:"bytes_rx"`
			BytesTx    int    `json:"bytes_tx"`
			PacketsRx  int    `json:"packets_rx"`
			PacketsTx  int    `json:"packets_tx"`
			ErrorsRx   int    `json:"errors_rx"`
			ErrorsTx   int    `json:"errors_tx"`
			DroppedRx  int    `json:"dropped_rx"`
			DroppedTx  int    `json:"dropped_tx"`
			Collisions int    `json:"collisions"`
		} `json:"wired"`
		WanStatus []struct {
			Interface     string `json:"interface"`
			LinkSpeedMbps int    `json:"link_speed_mbps"`
			CarrierLink   bool   `json:"carrier_link"`
			FullDuplex    bool   `json:"full_duplex,omitempty"`
		} `json:"wan_status"`
		CountryCodeConflict   int    `json:"country_code_conflict"`
		ConfiguredCountryName string `json:"configured_country_name"`
		ConfiguredCountryCode int    `json:"configured_country_code"`
		CliEnabled            bool   `json:"cli_enabled"`
		Region                string `json:"region"`
		Location              string `json:"location"`
		WtpMode               string `json:"wtp_mode"`
		LldpEnable            bool   `json:"lldp_enable"`
		Lldp                  []struct {
			LocalPort         string `json:"local_port"`
			ChassisId         string `json:"chassis_id"`
			SystemName        string `json:"system_name"`
			SystemDescription string `json:"system_description"`
			Capability        string `json:"capability"`
			PortId            string `json:"port_id"`
			PortDescription   string `json:"port_description"`
			MauOperatingMode  string `json:"mau_operating_mode"`
			Ip                string `json:"ip"`
			Vlan              int    `json:"vlan"`
		} `json:"lldp"`
		LedBlink        bool `json:"led_blink"`
		CpuUsage        int  `json:"cpu_usage"`
		MemFree         int  `json:"mem_free"`
		MemTotal        int  `json:"mem_total"`
		IsWpa3Supported bool `json:"is_wpa3_supported"`
	} `json:"results"`
	Vdom    string `json:"vdom"`
	Path    string `json:"path"`
	Name    string `json:"name"`
	Action  string `json:"action"`
	Status  string `json:"status"`
	Serial  string `json:"serial"`
	Version string `json:"version"`
	Build   int    `json:"build"`
}

func (maReq WifiManagedAPRequest) Body() []byte {
	return []byte{}
}

func (maRsp *WifiManagedAPResponse) Init(body io.ReadCloser) error {
	if b, err := ioutil.ReadAll(body); err != nil {
		return err
	} else {
		if err := json.Unmarshal(b, &maRsp); err != nil {
			return err
		}
	}
	return nil
}

func (maRsp *WifiManagedAPResponse) String() string {
	var aps string

	for _, ap := range maRsp.Results {
		aps += fmt.Sprintf("%s\t%s\t%d\n", ap.Name, ap.Status, ap.Clients)
	}
	return aps
}

func NewWifiManagedAP(api *GoFortiAPI.FortiAPI, serial string) *WifiManagedAP {
	// Creating userDeviceQuery
	wifiManagedAP := new(WifiManagedAP)
	wifiManagedAP.api = api

	// Setting URL to API
	wifiManagedAP.apiPath = "fgt/" + serial + "/api/v2/monitor/wifi/managed_ap"
	wifiManagedAP.request = []byte{}
	wifiManagedAP.response = new(WifiManagedAPResponse)

	return wifiManagedAP
}

func (wifi *WifiManagedAP) Post() error {
	return wifi.api.PostRequest(wifi.apiPath, wifi.request, wifi.response)
}

func (wifi *WifiManagedAP) Get() error {
	return wifi.api.GetRequest(wifi.apiPath, wifi.request, wifi.response)
}
