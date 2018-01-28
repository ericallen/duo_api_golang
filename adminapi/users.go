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
	Desc string `json:"desc,omitempty"`
	Name string `json:"name,omitempyt"`
}

type Phone struct {
	Phone_ID           string   `json:"phone_id,omitempty"`
	Number             string   `json:"number,omitempty"`
	Extention          string   `json:"extenstion,omitempty"`
	Name               string   `json:"name,omitempty"`
	Postdelay          string   `json:"postdelay,omitempty"`
	Predelay           string   `json:"predelay,omitempty"`
	Type               string   `json:"type,omitempty"`
	Capabilities       []string `json:"capabilities,omitempty"`
	Platform           string   `json:"platform,omitempty"`
	Activated          bool     `json:"activated,omitempty"`
	SMS_Passcodes_Sent bool     `json:"sms_passcodes_sent,omitempty"`
}

type Token struct {
	Serial   string `json:"serial,omitempty"`
	Token_ID string `json:"token_id,omitempty"`
	Type     string `json:"type,omitempty"`
}

type UserResponse struct {
	User_ID             string   `json:"user_id,omitempty"`
	Username            string   `json:"username,omitempty"`
	Alias1              string   `json:"alias1,omitempty"`
	Alias2              string   `json:"alias2,omitempty"`
	Alias3              string   `json:"alias3,omitempty"`
	Alias4              string   `json:"alias4,omitempty"`
	Created             int64    `json:"created,omitempty"`
	Firstname           string   `json:"firstname,omitempty"`
	Realname            string   `json:"realname,omitempty"`
	Email               string   `json:"email,omitempty"`
	Status              string   `json:"status,omitempty"`
	Groups              []Groups `json:"groups,omitempty"`
	Last_Directory_Sync int64    `json:"last_directory_sync,omitempty"`
	Last_Login          int64    `json:"last_login,omitempty"`
	Lastname            string   `json:"lastname,omitempty"`
	Notes               string   `json:"notes,omitempty"`
	Phones              []Phone  `json:"phones,omitempty"`
	Tokens              []Token  `json:"tokens,omitempty"`
}

type UsersResult struct {
	Stat     string         `json:"stat,omitempty"`
	Response []UserResponse `json:"response,omitempty"`
}

type UserResult struct {
	Stat     string       `json:"stat,omitempty"`
	Response UserResponse `json:"response,omitempty"`
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
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
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

func SetAlias1(alias1 string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("alias1", alias1)
	}
}

func SetAlias2(alias2 string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("alias2", alias2)
	}
}

func SetAlias3(alias3 string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("alias3", alias3)
	}
}

func SetAlias4(alias4 string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("alias4", alias4)
	}
}

func SetFirstname(firstname string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("firstname", firstname)
	}
}

func SetLastName(lastname string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("lastname", lastname)
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

//DeleteUser -
func (api *AdminApi) DeleteUser(user_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s", user_id)
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

//Enroll User
func (api *AdminApi) EnrollUser(username, email string, options ...func(*url.Values)) (*StatResult, error) {
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
	ret := &StatResult{}
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

type UserGroupResponse struct {
	Stat     string       `json:"stat,omitempty"`
	Response []UserGroups `json:"response,omitempty"`
}

type UserGroups struct {
	Desc     string `json:"desc,omitempty"`
	Group_ID string `json:"group_id,omitempty"`
	Name     string `json:"name,omitempty"`
}

//Retreive Groups by User ID
func (api *AdminApi) RetrieveGroupsByUserID(user_id string) (*UserGroupResponse, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/groups", user_id)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &UserGroupResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Assoicate Group with User
func (api *AdminApi) AssociateGroupwithUser(user_id string, options ...func(*url.Values)) (*StatResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}

	path := fmt.Sprintf("/admin/v1/users/%s/groups", user_id)
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

//Disassoicate Group from User
func (api *AdminApi) DisassociateGroupfromUser(user_id, group_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/groups/%s", user_id, group_id)
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

type PhoneResult struct {
	Stat     string  `json:"stat,omitempty"`
	Response []Phone `json:"response,omitempty"`
}

//Retreive Phones by User ID
func (api *AdminApi) RetreivePhonesbyUserID(user_id string) (*PhoneResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/phones", user_id)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &PhoneResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Disassociate Phone from User
func (api *AdminApi) DisassoicatePhonefromUser(user_id, phone_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/phones/%s", user_id, phone_id)
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

type TokenResult struct {
	Stat     string  `json:"stat,omitempty"`
	Response []Token `json:"response,omitempty"`
}

//Retreive Hardware Token by UserID
func (api *AdminApi) RetreiveHardwareTokenbyUserID(user_id string) (*TokenResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/tokens", user_id)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &TokenResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Associate Hardware Token with User
func (api *AdminApi) AssoicateHardwareTokenwithUser(user_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/tokens", user_id)
	_, body, err := api.SignedCall("POST", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &StatResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Disassociate Hardware Token from User
func (api *AdminApi) DisassociateHardwareTokenfromUser(user_id, token_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/tokens/%s", user_id, token_id)
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

//Retreive U2F Tokens by User ID
func (api *AdminApi) RetreiveU2FTokensbyUserID(user_id string) (*U2FTokenResponse, error) {
	path := fmt.Sprintf("/admin/v1/users/%s/u2ftokens", user_id)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &U2FTokenResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
