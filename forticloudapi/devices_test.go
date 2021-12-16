package forticloudapi

import (
	"github.com/get-code-ch/GoFortiAPI"
	"testing"
)

var ctx *GoFortiAPI.Context

func init() {
	ctx = GoFortiAPI.NewContext()
}

func TestGetDevice(t *testing.T) {
	var devices *DevicesGet

	devices = NewDevicesGet(ctx.FortiAPI, ctx.AccessToken)
	if devices == nil {
		t.Fatalf("Error creating DevicesGet object")
	}

	if err := devices.Get(); err != nil {
		t.Fatalf("Error getting device -> %v", err)
	}

	if count := len(*devices.Response); count < 1 {
		t.Fatalf("Fail: device length is %d should greater than 0", count)
	}

}
