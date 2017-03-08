package adminapi

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/ericallen/duo_api_golang"
)

type AuthenticationLogsResult struct {
	Stat     string
	Response []AuthenticationLog
}

type AuthenticationLog struct {
	Access_Device  AccessDevice
	Device         string
	Factor         string
	Integration    string
	IP             string
	Location       Location
	New_Enrollment bool
	Reason         string
	Result         string
	Timestamp      int64
	Username       string
}

type AccessDevice struct {
	Browser         string
	Browser_Version string
	Flash_Version   string
}

type Location struct {
	City    string
	State   string
	Country string
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
	Stat     string
	Response []AdminsitratorLog
}

type AdminsitratorLog struct {
	Action      string
	Description string
	Object      string
	Timestamp   int64
	Username    string
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
	Stat     string
	Response []TelephonyLog
}

type TelephonyLog struct {
	Timestamp int64
	Context   string
	Type      string
	Phone     string
	Credits   string
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
