package Tests

import (
	"github.com/get-code-ch/GoFortiAPI"
	"github.com/get-code-ch/GoFortiAPI/forticloudapi"
	"github.com/get-code-ch/SecretManager"
	"log"

	"testing"
)

var fortiAPI *GoFortiAPI.FortiAPI
var secret SecretManager.Secret

func init() {
	var err error

	vault := new(SecretManager.Vault)
	if err = vault.Open(); err != nil {
		log.Fatalf("Error getting secret -> %v", err)
	}
	defer vault.Close()

	if secret, err = vault.Read("Forti"); err != nil {
		log.Fatal("Error getting application login from Vault")
	}

}

func TestNewFortiAPI(t *testing.T) {

	t.Logf("Starting TestNewFortiAPI\n")
	region := secret.Parameters["region"]

	// Creating new FortiAPI instance
	fortiAPI = GoFortiAPI.NewFortiAPI(region)
	if fortiAPI.Region != region {
		t.Fatalf("Error region not correctly set fortiAPI.Region is %s should be %s", fortiAPI.Region, region)
	}

}

func TestNewAuthPost(t *testing.T) {

	t.Logf("Starting TestNewAuthPost\n")

	var err error

	var authPost *forticloudapi.AuthPost
	account := secret.Parameters["account"]
	username := secret.Username
	password := secret.Password

	// Getting API token and automatically renew it
	// Getting authentication from API
	if authPost, err = forticloudapi.NewAuthPost(fortiAPI, account, username, password); err != nil {
		t.Fatalf("Error setting authPost -> %v", err)
	}

	if err = authPost.Post(); err != nil {
		log.Fatalf("Error getting API Token -> %v", err)
	}

}
