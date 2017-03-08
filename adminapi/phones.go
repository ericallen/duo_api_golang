package adminapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ericallen/duo_api_golang"
)

type PhoneUserResult struct {
	Stat     string
	Response []PhoneResponse
}

type PhoneResponse struct {
	Activated          bool
	Capabilities       []string
	Encrypted          string
	Extension          string
	Fringerprint       string
	Name               string
	Number             string
	Phone_Id           string
	Platform           string
	Postdelay          string
	Predelay           string
	Screenlock         string
	SMS_Passcodes_Sent bool
	Type               string
	Users              []PhoneUser
}

type PhoneUser struct {
	Email      string
	Last_Login int64
	Notes      string
	Realname   string
	Status     string
	User_ID    string
	Username   string
}

type PhoneIDResult struct {
	Stat     string
	Response PhoneResponse
}

func SetNumber(number string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("number", number)
	}
}

func SetExtension(extension string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("extension", extension)
	}
}

//Retrieve Phones
//Optional parameters SetNumber and SetExtension
func (api *AdminApi) RetrievePhones() (*PhoneUserResult, error) {
	_, body, err := api.SignedCall("GET", "/admin/v1/phones", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &PhoneUserResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err
}

func SetName(name string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("name", name)
	}
}

func SetType(typestring string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("type", typestring)
	}
}

func SetPlatform(platform string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("platform", platform)
	}
}

func SetPredelay(predelay string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("predelay", predelay)
	}
}

func SetPostdelay(postdelay string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("postdelay", postdelay)
	}
}

//Create Phone
//Optional parameters SetNumber, SetName, SetExtension, SetType, SetPlatform, SetPredelay and SetPostdelay
func (api *AdminApi) CreatePhone(options ...func(*url.Values)) (*PhoneIDResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("POST", "/admin/v1/phones", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &PhoneIDResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Retrieve Phone by ID
//Required parameters - phone_id
//Optional parameters - None
func (api *AdminApi) RetrievePhoneByID(phone_id string) (*PhoneIDResult, error) {
	path := fmt.Sprintf("/admin/v1/phones/%s", phone_id)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &PhoneIDResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Modify Phone
//Required parameters - phone_id
//Optional parameters - SetNumber, SetName, SetExtension, SetType, SetPlatform, SetPredelay, SetPostdelay
func (api *AdminApi) ModifyPhone(phone_id string, options ...func(*url.Values)) (*PhoneIDResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	path := fmt.Sprintf("/admin/v1/phones/%s", phone_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &PhoneIDResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Delete Phone
//Required parameters - phone_id
//Optional parameters - None
func (api *AdminApi) DeletePhone(phone_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/phones/%s", phone_id)
	_, body, err := api.SignedCall("DELETE", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &StatResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err
}

func SetValidSeconds(valid_seconds string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("valid_secs", valid_seconds)
	}
}

func SetInstall(install string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("install", install)
	}
}

type ActivationCodeResponse struct {
	Activation_URL     string
	Activation_Barcode string
	Installation_URL   string
	Valid_Secs         int64
}

type ActivationCodeResult struct {
	Stat     string
	Response ActivationCodeResponse
}

//Create Activiation Code
//Required parameters - phone_id
//Optional parameters - SetValidSeconds, SetInstall
func (api *AdminApi) CreateActivationCode(phone_id string, options ...func(*url.Values)) (*ActivationCodeResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	path := fmt.Sprintf("/admin/v1/phones/%s/activation_url", phone_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &ActivationCodeResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err
}

type ActivationCodeSMSResult struct {
	Stat     string
	Response struct {
		Activation_Barcode string
		Activation_Msg     string
		Installation_Msg   string
		Valid_Secs         int64
	}
}

func SetInstallationMsg(installation_msg string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("install", installation_msg)
	}
}

func SetActivationMsg(activation_msg string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("install", activation_msg)
	}
}

//Send Activiation Code by SMS
//Required parameters - phone_id
//Optional parameters - SetValidSeconds, SetInstall, SetInstallationMsg, SetActivationMsg
func (api *AdminApi) SendActivationCodebySMS(phone_id string, options ...func(*url.Values)) (*ActivationCodeSMSResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	path := fmt.Sprintf("/admin/v1/phones/%s/send_sms_activation", phone_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &ActivationCodeSMSResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type InstallationURLviaSMSResult struct {
	Stat     string
	Response struct {
		Installation_Msg string
	}
}

//Send Installational URL vis SMS
//Required parameters - phone_id
//Optional parameters - SetInstallationMsg
func (api *AdminApi) SendInstallionalURLviaSMS(phone_id string, options ...func(*url.Values)) (*InstallationURLviaSMSResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	path := fmt.Sprintf("/admin/v1/phones/%s/send_sms_installation", phone_id)
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &InstallationURLviaSMSResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Send Passcodes via SMS
//Required parameters - phone_id
//Optional Parameters - None
func (api *AdminApi) SendPasscodesviaSMS(phone_id string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/phones/%s/send_sms_passcodes", phone_id)
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
