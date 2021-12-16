package fortiosmonitor

import (
	"github.com/get-code-ch/GoFortiAPI"
	"log"
	"testing"
)

var ctx *GoFortiAPI.Context

func init() {
	ctx = GoFortiAPI.NewContext()
}

func TestGetUser(t *testing.T) {
	var user *UserDeviceQuery

	user = NewUserDeviceQuery(ctx.FortiAPI, ctx.AccessToken, "FGT40FTK21032375")
	if user == nil {
		t.Fatalf("Error creating DevicesGet object")
	}

	if err := user.Get(); err != nil {
		t.Fatalf("Error getting device -> %v", err)
	}
	log.Printf("User -> %v", user)

}
