package adminapi

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/ericallen/duo_api_golang"
)

//Return object for the "Retrieve Summary" call
type SummaryResult struct {
	StatResult
	Response struct {
		Admin_Count                 int
		Integration_Count           int
		Telephony_Credits_Remaining int
		User_count                  int
	}
}

//RetrieveSummary
//Required parameters: none
//Optional parameters: none
func (api *AdminApi) RetrieveSummary() (*SummaryResult, error) {
	_, body, err := api.SignedCall("GET", "/admin/v1/info/summary", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &SummaryResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Return object for the "Telephony Credits Used Report" call
type TelephonyCreditsUsedResult struct {
	StatResult
	Response struct {
		Maxtime                int
		Mintime                int
		Telephony_Credits_Used int
	}
}

func AccountInfoMaxTime(maxtime int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("maxtime", strconv.Itoa(maxtime))
	}
}

func AccountInfoMinTime(mintime int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("mintime", strconv.Itoa(mintime))
	}
}

//Telephony Credits Used Report
//Required parameters: none
//Optional parameters: AccountInfoMaxTime, AccountInfoMinTime
func (api *AdminApi) TelephonyCreditsUsedReport(options ...func(*url.Values)) (*TelephonyCreditsUsedResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("GET", "/admin/v1/info/telephony_credits_used", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &TelephonyCreditsUsedResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Return object for the "Authentication Attempts Report" call
type AuthenticationAttemptsResult struct {
	StatResult
	Response struct {
		Maxtime                 int
		Mintime                 int
		Authentication_Attempts struct {
			Error   int
			Failure int
			Fraud   int
			Success int
		}
	}
}

//Authentication Attempts Report
//Required parameters: none
//Optional parameters: AccountInfoMaxTime, AccountInfoMinTime
func (api *AdminApi) AuthenticationAttemptsReport(options ...func(*url.Values)) (*AuthenticationAttemptsResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("GET", "/admin/v1/info/authentication_attempts", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &AuthenticationAttemptsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Retun object for the "Users with Authentication Attempts Report" call
type UsersAuthenticationAttemptsResult struct {
	StatResult
	Response struct {
		Mintime                    int
		Maxtime                    int
		UserAuthenticationAttempts struct {
			Error   int
			Failure int
			Fraud   int
			Success int
		}
	}
}

//Users with Authentication Attempts Report
//Required parameters: none
//Optional parameters: AccountInfoMaxTime, AccountInfoMinTime
func (api *AdminApi) UserWithAuthenticationAttemptsReport(options ...func(*url.Values)) (*UsersAuthenticationAttemptsResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("GET", "/admin/v1/info/user_authentication_attempts", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &UsersAuthenticationAttemptsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, err
}
