package GoFortiAPI

import (
	"bytes"
	"io"
	"net/http"
)

// AccessToken Authentication Token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type FortiAPI struct {
	Region      string
	AccessToken AccessToken
	HttpClient  *http.Client
}

type APICall struct {
	Header http.Header
	URL    string
	Body   []byte
}

// Response interface definition
type Response interface {
	Init(io.ReadCloser) error
	String() string
}

type Request interface {
	Post() error
	Get() error
	API() APICall
}

// NewFortiAPI function create new FortiAPI object for accessing FortiGate from FortiCloud
func NewFortiAPI(region string) *FortiAPI {
	api := new(FortiAPI)
	api.HttpClient = &http.Client{}
	api.AccessToken = AccessToken{}
	api.Region = region

	return api
}

//GetRequest send a Get Request to FortiGate Cloud
func (api *FortiAPI) GetRequest(request Request, response Response) error {
	apiCall := request.API()

	httpRequest, _ := http.NewRequest("GET", apiCall.URL, bytes.NewReader(apiCall.Body))
	httpRequest.Header = apiCall.Header

	if httpResponse, err := api.HttpClient.Do(httpRequest); err != nil {
		return err
	} else {
		return response.Init(httpResponse.Body)
	}
}

//PostRequest send a Post Request to FortiGate Cloud
func (api *FortiAPI) PostRequest(request Request, response Response) error {
	apiCall := request.API()

	httpRequest, _ := http.NewRequest("POST", apiCall.URL, bytes.NewReader(apiCall.Body))
	httpRequest.Header = apiCall.Header

	if httpResponse, err := api.HttpClient.Do(httpRequest); err != nil {
		return err
	} else {
		return response.Init(httpResponse.Body)
	}
}
