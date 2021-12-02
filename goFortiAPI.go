package GoFortiAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// AccessToken Authentication Token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// AuthPostBody Authentication POST request Body type
type AuthPostBody struct {
	AccountId string `json:"accountId"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}

type FortiAPI struct {
	AccessToken AccessToken
	HttpClient  *http.Client
}

// Response interface definition
type Response interface {
	Init(io.ReadCloser) error
	String() string
}

// NewCloudFortiAPI function create new FortiAPI object for accessing FortiGate from FortiCloud
func NewCloudFortiAPI(region string, account string, username string, password string) (*FortiAPI, error) {
	// Initializing properties
	api := new(FortiAPI)
	api.HttpClient = &http.Client{}
	api.AccessToken = AccessToken{}

	// Getting AccessToken from FortiGate Cloud
	authUrl := fmt.Sprintf("https://%s/forticloudapi/v1/auth", region)
	authBody, _ := json.Marshal(AuthPostBody{account, username, password})

	request, _ := http.NewRequest("POST", authUrl, bytes.NewReader(authBody))
	request.Header.Set("Content-Type", "application/json")

	return api, nil
}

//GetRequest send a Get Request to FortiGate Cloud or directly to FortiGate
func (api *FortiAPI) GetRequest(url string, response Response, accessToken string) (Response, error) {

	request, _ := http.NewRequest("GET", url, bytes.NewReader(nil))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	if httpResponse, err := api.HttpClient.Do(request); err != nil {
		return nil, err
	} else {
		err := response.Init(httpResponse.Body)
		return response, err
	}
}

//PostRequest send a Get Request to FortiGate Cloud or directly to FortiGate
func PostRequest(url string, response Response, accessToken string) (Response, error) {

	request, _ := http.NewRequest("GET", url, bytes.NewReader(nil))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := &http.Client{}

	if httpResponse, err := client.Do(request); err != nil {
		return nil, err
	} else {
		err := response.Init(httpResponse.Body)
		return response, err
	}
}
