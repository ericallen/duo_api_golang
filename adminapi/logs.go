package adminapi

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/ericallen/duo_api_golang"
)

type AuthenticationLogsResult struct {
	Stat     string              `json:"stat,omitempty"`
	Response []AuthenticationLog `json:"response,omitempty"`
}

type AuthenticationLog struct {
	Access_Device  AccessDevice `json:"access_device,omitempty"`
	Device         string       `json:"device,omitempty"`
	Factor         string       `json:"factor,omitempty"`
	Integration    string       `json:"integration,omitempty"`
	IP             string       `json:"ip,omitempty"`
	Location       Location     `json:"location,omitempty"`
	New_Enrollment bool         `json:"new_enrollment,omitempty"`
	Reason         string       `json:"reason,omitempty"`
	Result         string       `json:"result,omitempty"`
	Timestamp      int64        `json:"timestamp,omitempty"`
	Username       string       `json:"username,omitempty"`
}

type AccessDevice struct {
	Browser                 string `json:"browser,omitempty"`
	Browser_Version         string `json:"browser_version,omitempty"`
	Flash_Version           string `json:"flash_version,omitempty"`
	Java_Version            string `json:"java_version,omitempty"`
	OS                      string `json:"os,omitempty"`
	OS_Version              string `json:"os_version,omitempty"`
	Trusted_Endpoint_Status string `json:"trusted_endpoint_status,omitempty"`
}

type Location struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}

func LogsMinTime(mintime int64) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("mintime", strconv.FormatInt(mintime, 10))
	}
}

//Authentication Logs
//Required parameters: none
//Optional parameters: LogsMinTime
func (api *AdminApi) AuthenticationLogs(options ...func(*url.Values)) (*AuthenticationLogsResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("GET", "/admin/v1/logs/authentication", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &AuthenticationLogsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type AdministratorLogResult struct {
	Stat     string             `json:"stat,omitempty"`
	Response []AdminsitratorLog `json:"response,omitempty"`
}

type AdminsitratorLog struct {
	Action      string `json:"action,omitempty"`
	Description string `json:"description,omitempty"`
	Object      string `json:"object,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	Username    string `json:"username,omitempty"`
}

//Administrator Logs
//Required parameters: none
//Optional parameters: LogsMinTime
func (api *AdminApi) AdministratorLogs(options ...func(*url.Values)) (*AdministratorLogResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("GET", "/admin/v1/logs/administrator", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &AdministratorLogResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type TelephonyLogResult struct {
	Stat     string         `json:"stat,omitempty"`
	Response []TelephonyLog `json:"response,omitempty"`
}

type TelephonyLog struct {
	Timestamp int64  `json:"timestamp,omitempty"`
	Context   string `json:"context,omitempty"`
	Type      string `json:"type,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Credits   int64  `json:"credits,omitempty"`
}

//Telephony Logs
//Required parameters: none
//Optional parameters: LogsMinTime
func (api *AdminApi) TelephonyLogs(options ...func(*url.Values)) (*TelephonyLogResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("GET", "/admin/v1/logs/telephony", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &TelephonyLogResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
