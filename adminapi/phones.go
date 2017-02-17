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

func (api *AdminApi) RetrievePhones() (*PhoneUserResult, error) {
	opts := url.Values{}

	_, body, err := api.SignedCall("GET", "/admin/v1/phones", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &PhoneUserResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err
}

func (api *AdminApi) DeletePhone(phone_id string) (*SimpleResponse, error) {
	opts := url.Values{}
	path := fmt.Sprintf("/admin/v1/phones/%s", phone_id)
	_, body, err := api.SignedCall("DELETE", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &SimpleResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err
}
