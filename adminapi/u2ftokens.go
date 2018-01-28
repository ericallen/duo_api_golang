package adminapi

import (
	"encoding/json"
	"fmt"

	"github.com/ericallen/duo_api_golang"
)

type U2FToken struct {
	Date_Added      int64  `json:"date_added,omitempty"`
	Registration_ID string `json:"registration_id,omitempty`
	User            User   `json:"user,omitempty`
}

type U2FTokenResponse struct {
	Stat     string     `json:"stat,omitempty`
	Response []U2FToken `json:"response,omitempty"`
}

func (api *AdminApi) RetrieveU2FTokens() (*U2FTokenResponse, error) {
	_, body, err := api.SignedCall("GET", "/admin/v1/u2ftokens", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &U2FTokenResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (api *AdminApi) RetrieveU2FTokenByID(registration_id string) (*U2FTokenResponse, error) {
	path := fmt.Sprintf("/admin/v1/u2ftokens/%s", registration_id)
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

//Disassoicate Group from User
func (api *AdminApi) DeleteU2FToken(registration_id string) (*U2FTokenResponse, error) {
	path := fmt.Sprintf("/admin/v1/u2ftokens/%s", registration_id)
	_, body, err := api.SignedCall("DELETE", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &U2FTokenResponse{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
