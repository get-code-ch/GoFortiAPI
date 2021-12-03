package forticloudapi

import (
	"encoding/json"
	"fmt"
	"github.com/get-code-ch/GoFortiAPI"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AuthPostBody struct {
	AccountId string `json:"accountId"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}

// AuthPost Authentication POST request Body type
type AuthPost struct {
	api      *GoFortiAPI.FortiAPI
	url      url.URL
	header   http.Header
	body     []byte
	Response *AuthResponse
}

type AuthResponse struct {
	AccessToken GoFortiAPI.AccessToken
}

func NewAuthResponse() *AuthResponse {
	authResponse := new(AuthResponse)
	return authResponse
}

func (r *AuthResponse) Init(body io.ReadCloser) error {
	if b, err := ioutil.ReadAll(body); err != nil {
		return err
	} else {
		if err := json.Unmarshal(b, &r.AccessToken); err != nil {
			return err
		}
	}
	return nil
}

func (r *AuthResponse) String() string {
	var response string

	response = fmt.Sprintf("Token:%s, ExpiresIn:%d", r.AccessToken.AccessToken, r.AccessToken.ExpiresIn)
	return response
}

func NewAuthPost(api *GoFortiAPI.FortiAPI, account string, username string, password string) (*AuthPost, error) {
	//
	var err error

	// Creating authPost
	authPost := new(AuthPost)

	// Setting URL to API
	authPost.url.Scheme = "https"
	authPost.url.Host = api.Region
	authPost.url.Path = "forticloudapi/v1/auth"

	// Setting Header
	authPost.header = http.Header{}
	authPost.header.Set("Content-Type", "application/json")
	authPost.Response = NewAuthResponse()

	// Setting AuthPost Body
	if authPost.body, err = json.Marshal(AuthPostBody{account, username, password}); err != nil {
		return nil, err
	}
	authPost.api = api

	return authPost, nil
}

func (request *AuthPost) API() GoFortiAPI.APICall {
	return GoFortiAPI.APICall{Header: request.header, URL: request.url.String(), Body: request.body}
}

func (request *AuthPost) Post() error {
	return request.api.PostRequest(request, request.Response)
}

func (request *AuthPost) Get() error {
	return request.api.GetRequest(request, request.Response)
}
