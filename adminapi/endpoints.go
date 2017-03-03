package adminapi

import (
	"encoding/json"
	"fmt"

	"github.com/ericallen/duo_api_golang"
)

type EndpointsResult struct {
	Stat     string
	Response []Endpoint
}

type Endpoint struct {
	Browsers   []Browser
	Email      string
	Epkey      string
	Model      string
	OS_Family  string
	OS_Version string
	Type       string
	Username   string
}

type Browser struct {
	Browser_Family  string
	Browser_Version string
	Flash_Version   string
	Java_Version    string
	Last_Used       int64
}

//RetrieveEndpoints
//Required parameters: None
//Optional parameters: None
func (api *AdminApi) RetrieveEndpoints() (*EndpointsResult, error) {
	_, body, err := api.SignedCall("GET", "/admin/v1/endpoints", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &EndpointsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Retrieve Endpoints by ID
//Required parameters: epkey
//Optional parameters: None
func (api *AdminApi) RetrieveEndpointsbyID(epkey string) (*EndpointsResult, error) {
	path := fmt.Sprintf("/admin/v1/endpoints/%s", epkey)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &EndpointsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
