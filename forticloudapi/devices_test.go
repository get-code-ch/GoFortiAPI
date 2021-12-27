package forticloudapi

import (
	"github.com/get-code-ch/GoFortiAPI"
	"github.com/get-code-ch/SecretManager"
	"log"
	"testing"
)

func TestNewDevicesList(t *testing.T) {
	var err error
	var secret SecretManager.Secret
	var api *GoFortiAPI.FortiAPI
	var devices *Devices

	// Getting parameters from secret vault
	vault := new(SecretManager.Vault)
	if err := vault.Open(); err != nil {
		log.Fatalf("Error getting secret -> %v", err)
	}
	defer vault.Close()

	if secret, err = vault.Read("Forti"); err != nil {
		log.Fatal("Error getting application login from Vault")
	}

	account := secret.Parameters["account"]
	region := secret.Parameters["region"]
	username := secret.Username
	password := secret.Password

	api, err = GoFortiAPI.NewFortiAPI(region, account, username, password)

	if err != nil {
		t.Fatalf("Error creating FortiAPI -> %v", err)
	}

	devices = NewDevicesList(api)
	devices.Get()
	log.Printf("%v", devices.response)
}
