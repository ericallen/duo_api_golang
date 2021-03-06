package adminapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ericallen/duo_api_golang"
)

type GroupResponse struct {
	Desc               string `json:"desc,omitemtpy"`
	Push_Enabled       bool   `json:"push_enabled,omitempty"`
	SMS_Enabled        bool   `json:"sms_enabled,omitempty"`
	Voice_Enabled      bool   `json:"voice_enabled,omitempty"`
	Mobile_OTP_Enabled bool   `json:"mobile_otp_enabled,omitempty"`
	Group_ID           string `json:"group_id,omitemtpy"`
	Name               string `json:"name,omitempty"`
	Status             string `json:"status,omitempty"`
}

type GroupsResult struct {
	Stat     string          `json:"stat,omitempty"`
	Response []GroupResponse `json:"response,omitempty"`
}

type GroupResult struct {
	Stat     string        `json:"stat,omitempty"`
	Response GroupResponse `json:"response,omitempty"`
}

//RetrieveGroups
//Required parameters - None
//Optional parameters - None
func (api *AdminApi) RetrieveGroups() (*GroupsResult, error) {
	_, body, err := api.SignedCall("GET", "/admin/v1/groups", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &GroupsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err
}

func GroupDesc(desc string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("desc", desc)
	}
}

func GroupPushEnabled(push_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("push_enabled", strconv.FormatBool(push_enabled))
	}
}

func GroupSMSEnabled(sms_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("sms_enabled", strconv.FormatBool(sms_enabled))
	}
}

func GroupVoiceEnabled(voice_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("voice_enabled", strconv.FormatBool(voice_enabled))
	}
}

func GroupMobileOTPEnabled(mobile_otp_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("mobile_otp_enabled", strconv.FormatBool(mobile_otp_enabled))
	}
}

func GroupStatus(status string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("status", status)
	}
}

//CreateGroup
//Required parameters - name
//Optional parameters - GroupDesc, GroupPushEnabled, GroupSMSEnabled, GroupVoiceEnabled, GroupMobileOTPEnabled, GroupStatus
func (api *AdminApi) CreateGroup(name string, options ...func(*url.Values)) (*GroupResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("POST", "/admin/v1/groups", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &GroupResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//GetGroupInfo
//Required parameters - group_id
//Optional parameters - none
func (api *AdminApi) GetGroupInfo(group_id string) (*GroupResult, error) {
	path := fmt.Sprintf("/admin/v1/groups/%s", group_id)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &GroupResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func GroupName(name string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("name", name)
	}
}

//UpdateGroup
//Required parameters - group_id
//Optional parameters - GroupName, GroupDesc, GroupPushEnabled, GroupSMSEnabled, GroupVoiceEnabled, GroupMobileOTPEnabled, GroupStatus
func (api *AdminApi) UpdateGroup(group_id string, options ...func(*url.Values)) (*GroupResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	path := fmt.Sprintf("/admin/v1/groups/%s", group_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &GroupResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//DeleteGroup
//Required parameters - group_id
//Optional parameters - none
func (api *AdminApi) DeleteGroup(group_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/groups/%s", group_id)
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
