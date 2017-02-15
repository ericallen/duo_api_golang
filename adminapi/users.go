package adminapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ericallen/duo_api_golang"
)

// Optional parameter for the Retreiveusers method.
func SetUsername(username string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("username", username)
	}
}

type Groups struct {
	Desc string
	Name string
}

type Phones struct {
	Phone_ID           string
	Number             string
	Extention          string
	Name               string
	Postdelay          string
	Predelay           string
	Type               string
	Capabilities       []string
	Platform           string
	Activated          bool
	SMS_Passcodes_Sent bool
}

type Tokens struct {
	Serial   string
	Token_ID string
	Type     string
}

type UserResponse struct {
	User_ID    string
	Username   string
	Realname   string
	Email      string
	Status     string
	Groups     []Groups
	Last_Login int64
	Notes      string
	Phones     []Phones
	Tokens     []Tokens
}

type UsersResult struct {
	Stat     string
	Response []UserResponse
}

type UserResult struct {
	Stat     string
	Response UserResponse
}

// Users - Retreive Users
func (api *AdminApi) RetrieveUsers(options ...func(*url.Values)) (*UsersResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}

	_, body, err := api.SignedCall("GET", "/admin/v1/users", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &UsersResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Users - Create Users
func (api *AdminApi) CreateUser(username string, options ...func(*url.Values)) (*UserResult, error) {
	opts := url.Values{}
	opts.Set("username", username)
	for _, o := range options {
		o(&opts)
	}

	_, body, err := api.SignedCall("POST", "/admin/v1/users", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &UserResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err

}

// Users - Retrive User by ID
func (api *AdminApi) RetrieveUserByID(userid string) (*UserResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s", userid)
	opts := url.Values{}
	_, body, err := api.SignedCall("GET", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}

	ret := &UserResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func SetRealname(realname string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("realname", realname)
	}
}

func SetEmail(email string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("email", email)
	}
}

func SetStatus(status string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("status", status)
	}
}

func SetNotes(notes string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("notes", notes)
	}
}

//ModifyUser -
func (api *AdminApi) ModifyUser(user_id string, options ...func(*url.Values)) (*UserResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}

	path := fmt.Sprintf("/admin/v1/users/%s", user_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &UserResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type SimpleResponse struct {
	Stat     string
	Response string
}

//DeleteUser -
func (api *AdminApi) DeleteUser(user_id string) (*SimpleResponse, error) {
	opts := url.Values{}
	path := fmt.Sprintf("/admin/v1/users/%s", user_id)
	_, body, err := api.SignedCall("DELETE", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &SimpleResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Enroll User
func (api *AdminApi) EnrollUser(username, email string, options ...func(*url.Values)) (*SimpleResponse, error) {
	opts := url.Values{}
	opts.Set("username", username)
	opts.Set("email", email)
	for _, o := range options {
		o(&opts)
	}

	_, body, err := api.SignedCall("POST", "/admin/v1/users/enroll", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &SimpleResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type BypassCodesResponse struct {
	Stat     string
	Response []string
}

func SetBypassCount(count string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("count", count)
	}
}

func SetBypassCodes(codes string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("codes", codes)
	}
}

func SetBypassReuseCount(reuse_count string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("reuse_count", reuse_count)
	}
}

//Get Bypass Codes for User
func (api *AdminApi) GetBypassCodesForUser(user_id string, options ...func(*url.Values)) (*BypassCodesResponse, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}

	path := fmt.Sprintf("/admin/v1/users/%s/bypass_codes", user_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &BypassCodesResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
