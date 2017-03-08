package adminapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ericallen/duo_api_golang"
)

type TokensResult struct {
	Stat     string
	Response []TokenResponse
}

type TokenResponse struct {
	Serial    string
	Token_ID  string
	Type      string
	Totp_Step string
	Users     []User
}

type User struct {
	User_id    string
	Username   string
	Realname   string
	Email      string
	Status     string
	Last_Login int64
	Notes      string
}

func TokenType(tokentype string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("type", tokentype)
	}
}

func TokenSerial(tokenserial string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("serial", tokenserial)
	}
}

//Retrieve Hardware Token
//Required parameters - none
//Optional parameters - TokenType, TokenSerial
func (api *AdminApi) RetrieveHardwareToken(options ...func(*url.Values)) (*TokensResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("GET", "/admin/v1/tokens", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &TokensResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func TokenSecret(secret string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("secret", secret)
	}
}

func TokenCounter(counter int64) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("counter", strconv.FormatInt(counter, 10))
	}
}

func TokenPrivateId(private_id string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("private_id", private_id)
	}
}

func TokenAesKey(aes_key string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("aes_key", aes_key)
	}
}

//Create Hardware Token
//Required parameters - TokenType, TokenSerial
//Option parameters - TokenSecret, TokenCounter, TokenPrivateId, TokenAesKey
func (api *AdminApi) CreateHardwareToken(tokentype, tokenserial string, options ...func(*url.Values)) (*TokenResult, error) {
	opts := url.Values{}
	opts.Set("type", tokentype)
	opts.Set("serial", tokenserial)
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("POST", "/admin/v1/tokens", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &TokenResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Retrieve Hardware Token by ID
//Required parameters - token_id
//Option parameters - none
func (api *AdminApi) RetrieveHardwareTokenbyID(token_id string) (*TokensResult, error) {
	path := fmt.Sprintf("/admin/v1/tokens/%s", token_id)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &TokensResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Resync Hardware Token
//Required parameters - token_id, code1, code2, code3
//Optional parameters - none
func (api *AdminApi) ResyncHardwareToken(token_id, code1, code2, code3 string) (*StatResult, error) {
	opts := url.Values{}
	opts.Set("code1", code1)
	opts.Set("code2", code2)
	opts.Set("code3", code3)
	path := fmt.Sprintf("/admin/v1/tokens/%s/resync", token_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &StatResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Delete Hardware Token
//Required parameters - token_id
//Optional parameters - none
func (api *AdminApi) DeleteHardwareToken(token_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/tokens/%s", token_id)
	_, body, err := api.SignedCall("DELETE", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &StatResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
