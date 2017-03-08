package adminapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/ericallen/duo_api_golang"
)

type IntegrationsResult struct {
	Stat     string
	Response []Integration
}

type IntegrationResult struct {
	Stat     string
	Response Integration
}

type Integration struct {
	Enroll_Policy                 string
	Greeting                      string
	Groups_Allowed                []GroupResponse
	Integration_Key               string
	IP_Whitelist                  []string
	IP_Whitelist_Enroll_Policy    string
	Name                          string
	Notes                         string
	Secret_Key                    string
	Type                          string
	Trusted_Device_Days           int
	Username_Normalization_Policy string
}

//Retrieve Integrations
//Required parameters - none
//Optional parameters - none
func (api *AdminApi) RetrieveIntegrations() (*IntegrationsResult, error) {
	_, body, err := api.SignedCall("GET", "/admin/v1/integrations", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &IntegrationsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Retrieve Integration by key
//Required parameters - integration_key
//Optional parameters - none
func (api *AdminApi) RetrieveIntegrationbyKey(integration_key string) (*IntegrationResult, error) {
	path := fmt.Sprintf("/admin/v1/integration/%s", integration_key)
	_, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &IntegrationResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func IntegrationEnrollPolicy(enroll_policy string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("enroll_policy", enroll_policy)
	}
}

func IntegrationGreeting(greeting string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("greeting", greeting)
	}
}

func IntegrationGroupsAllowed(groups_allowed []string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("groups_allowed", strings.Join(groups_allowed, ","))
	}
}

func IntegrationNotes(notes string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("notes", notes)
	}
}

func IntegrationAdminApiAdmins(adminapi_admins bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("adminapi_admins", Btoa(adminapi_admins))
	}
}

func IntegrationAdminApiInfo(adminapi_info bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("adminapi_info", Btoa(adminapi_info))
	}
}

func IntegrationAdminApiIntegrations(adminapi_integrations bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("adminapi_integrations", Btoa(adminapi_integrations))
	}
}

func IntegrationAdminApiReadLog(adminapi_read_log bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("adminapi_read_log", Btoa(adminapi_read_log))
	}
}

func IntegrationAdminApiReadResource(adminapi_read_resource bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("adminapi_read_resources", Btoa(adminapi_read_resource))
	}
}

func IntegrationAdminApiSettings(adminapi_settings bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("adminapi_settings", Btoa(adminapi_settings))
	}
}

func IntegrationAdminApiWriteResource(adminapi_write_resource bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("adminapi_write_resource", Btoa(adminapi_write_resource))
	}
}

func IntegrationTrustedDeviceDays(trusted_device_days int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("trusted_device_days", strconv.Itoa(trusted_device_days))
	}
}

func IntegrationIPWhitelist(ip_whitelist string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("ip_whitelist", ip_whitelist)
	}
}

func IntegrationIPWhitelistEnrollPolicy(ip_whitelist_enroll_policy string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("ip_whitelist_enroll_policy", ip_whitelist_enroll_policy)
	}
}

func IntegrationUsernameNormalizationPolicy(username_normalization_policy string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("username_normalization_policy", username_normalization_policy)
	}
}

func IntegrationSelfServiceAllowed(self_service_allowed bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("self_service_allowed", Btoa(self_service_allowed))
	}
}

//Create Integration
//Required parameters - IntegrationName, IntegrationType
//Optional parameters - IntegrationEnrollPolicy, IntegrationGreeting, IntegrationGroupsAllowed, IntegrationNotes, IntegrationAdminapiAdmins, IntegrationAdminApiInfo, IntegrationAdminApiIntegrations, IntegrationAdminApiReadLog, IntegrationAdminApiReadResource, IntegrationAdminApiSettings, IntegrationAdminApiWriteResource, IntegrationTrustedDeviceDays, IntegrationIPWhitelist, IntegrationIPWhitelistEnrollPolicy, IntegrationUsernameNormalizationPolicy, IntegrationSelfServiceAllowed
func (api *AdminApi) CreateIntegration(IntegrationName, IntegrationType string, options ...func(*url.Values)) (*IntegrationsResult, error) {
	opts := url.Values{}
	opts.Set("name", IntegrationName)
	opts.Set("type", IntegrationType)
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("POST", "/admin/v1/integrations", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &IntegrationsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func IntegrationResetSecretKey(IntegrationResetSecretKey bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("reset_secret_key", Btoa(IntegrationResetSecretKey))
	}
}

//Modify Integration
//Required parameters - integration_key
//Optional parameters - IntegrationName, IntegrationEnrollPolicy, IntegrationGreeting, IntegrationGroupsAllowed, IntegrationNotes, IntegrationResetSecretKey, IntegrationAdminapiAdmins,  IntegrationAdminApiInfo, IntegrationAdminApiIntegrations, IntegrationAdminApiReadLog, IntegrationAdminApiReadResource, IntegrationAdminApiSettings, IntegrationAdminApiSettings, IntegrationAdminApiWriteResource, IntegrationTrustedDeviceDays, IntegrationIPWhitelist, IntegrationIPWhitelistEnrollPolicy, IntegrationUsernameNormalizationPolicy, IntegrationSelfServiceAllowed
func (api *AdminApi) ModifyIntegration(integration_key string, options ...func(*url.Values)) (*IntegrationResult, error) {
	path := fmt.Sprintf("/admin/v1/integration/%s", integration_key)
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &IntegrationResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Delete Integration
//Required parameters - integration_key
//Optional parameters - none
func (api *AdminApi) DeleteIntegration(integration_key string) (*StatResult, error) {
	path := fmt.Sprintf("/admin/v1/integration/%s", integration_key)
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

func Btoa(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
