package forticloudapi

import (
	"github.com/get-code-ch/GoFortiAPI"
	"github.com/get-code-ch/SecretManager"
	"log"
)

type context struct {
	AccessToken GoFortiAPI.AccessToken
	FortiAPI    *GoFortiAPI.FortiAPI
}

func NewContext() *context {
	var err error
	var secret SecretManager.Secret
	var authPost *AuthPost

	vault := new(SecretManager.Vault)
	if err = vault.Open(); err != nil {
		log.Fatalf("Error getting secret -> %v", err)
	}
	defer vault.Close()

	if secret, err = vault.Read("Forti"); err != nil {
		log.Fatal("Error getting application login from Vault")
	}

	region := secret.Parameters["region"]
	account := secret.Parameters["account"]
	username := secret.Username
	password := secret.Password

	// Creating new FortiAPI instance
	ctx := new(context)
	ctx.FortiAPI = GoFortiAPI.NewFortiAPI(region)

	// Getting API token and automatically renew it
	// Getting authentication from API
	if authPost, err = NewAuthPost(ctx.FortiAPI, account, username, password); err != nil {
		log.Fatalf("Error setting authPost -> %v", err)
	}

	if err = authPost.Post(); err != nil {
		log.Fatalf("Error getting API Token -> %v", err)
	}
	ctx.AccessToken = authPost.Response.AccessToken
	return ctx
}
