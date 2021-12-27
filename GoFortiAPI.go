package GoFortiAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

/*** Private data type ***/
// accessToken Authentication Token
type accessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// authBody body content for auth POST request
type authBody struct {
	AccountId string `json:"accountId"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}

/*** Public data type ***/

type FortiAPI struct {
	accessToken accessToken
	httpClient  *http.Client
	header      http.Header
	baseUrl     url.URL
}

// Response interface definition
type Response interface {
	Init(io.ReadCloser) error
	String() string
}

type Request interface {
	Body() []byte
}

// NewFortiAPI function create new FortiAPI object for accessing FortiGate from FortiCloud
func NewFortiAPI(region string, account string, username string, password string) (*FortiAPI, error) {
	// Creating FortiAPI object
	api := new(FortiAPI)

	api.baseUrl.Scheme = "https"
	api.baseUrl.Host = region
	api.baseUrl.Path = "forticloudapi/v1/"

	api.httpClient = &http.Client{}
	api.accessToken = accessToken{}

	// Setting header
	api.header = http.Header{}
	api.header.Set("Content-Type", "application/json")

	// Getting auth token
	if err := api.getAccessToken(authBody{AccountId: account, UserName: username, Password: password}); err != nil {
		return nil, err
	}

	// Launching access token automatic renewing every 6 hours
	tokenTicker := time.NewTicker(6 * time.Hour)
	go func() {
		for {
			<-tokenTicker.C
			if err := api.getAccessToken(authBody{AccountId: account, UserName: username, Password: password}); err != nil {
				log.Printf("Error getting API Token -> %v", err)
			}
		}
	}()

	return api, nil
}

// getAccessToken is a private method to get FortiAPI access token
func (api *FortiAPI) getAccessToken(authBody authBody) error {
	apiPath := api.baseUrl.String() + "auth"

	if body, err := json.Marshal(authBody); err != nil {
		return err
	} else {
		httpRequest, _ := http.NewRequest("POST", apiPath, bytes.NewReader(body))
		httpRequest.Header = api.header
		if httpResponse, err := api.httpClient.Do(httpRequest); err != nil {
			return err
		} else {
			if b, err := ioutil.ReadAll(httpResponse.Body); err != nil {
				return err
			} else {
				if err := json.Unmarshal(b, &api.accessToken); err != nil {
					return err
				}
			}
		}
	}
	// Adding Authorization header with accessToken value
	api.header.Set("Authorization", fmt.Sprintf("Bearer %s", api.accessToken.AccessToken))
	return nil
}

//GetRequest send a Get Request to FortiGate Cloud
func (api *FortiAPI) GetRequest(path string, request Request, response Response) error {
	apiPath := api.baseUrl.String() + path

	httpRequest, _ := http.NewRequest("GET", apiPath, bytes.NewReader(request.Body()))
	httpRequest.Header = api.header

	if httpResponse, err := api.httpClient.Do(httpRequest); err != nil {
		return err
	} else {
		return response.Init(httpResponse.Body)
	}
}

//PostRequest send a Post Request to FortiGate Cloud
func (api *FortiAPI) PostRequest(path string, request Request, response Response) error {
	apiPath := api.baseUrl.String() + path

	httpRequest, _ := http.NewRequest("POST", apiPath, bytes.NewReader(request.Body()))
	httpRequest.Header = api.header

	if httpResponse, err := api.httpClient.Do(httpRequest); err != nil {
		return err
	} else {
		return response.Init(httpResponse.Body)
	}
}
