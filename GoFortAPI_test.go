package GoFortiAPI

import (
	"github.com/get-code-ch/SecretManager"
	"log"
	"testing"
	"time"
)

func TestNewFortiAPI(t *testing.T) {
	var err error
	var secret SecretManager.Secret
	var api *FortiAPI

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

	api, err = NewFortiAPI(region, account, username, password)

	if err != nil {
		t.Fatalf("Error creating FortiAPI -> %v", err)
	}

	if api.accessToken.AccessToken == "" {
		t.Fatalf("Invalid access token")
	} else {
		log.Printf("token: %s - expired: %s", api.accessToken.AccessToken, time.Unix(api.accessToken.ExpiresIn, 0).Format(time.RFC822Z))
	}

}
