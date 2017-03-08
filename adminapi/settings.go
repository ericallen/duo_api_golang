package adminapi

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ericallen/duo_api_golang"
)

type SettingsResult struct {
	Stat     string
	Response Settings
}

type Settings struct {
	Caller_ID                     string
	Fraud_Email                   string
	Fraud_Email_Enabled           bool
	Inactive_User_Expiration      int
	Keypress_Confirm              string
	Keypress_Fraud                string
	Langauge                      string
	Lockout_Expire_Duration       int
	Lockout_Threshold             int
	Log_Retention_Days            int
	Minimum_Password_Length       int
	Mobile_OTP_Enabled            bool
	Name                          string
	Password_Required_Lower_Alpha bool
	Password_Required_Numeric     bool
	Password_Required_Special     bool
	Password_Requires_Upper_Alpha bool
	Push_Enabled                  bool
	SMS_Batch                     int
	SMS_Enabled                   bool
	SMS_Expiration                int
	SMS_Messages                  string
	SMS_Refresh                   int
	Telephony_Warning_Min         int
	Timezone                      string
	User_Telephony_Cost_Max       float32
	Voice_Enabled                 bool
}

//RetreiveSettings
//Require parameters: none
//Optional parameters: none
func (api *AdminApi) RetreiveSettings() (*SettingsResult, error) {
	_, body, err := api.SignedCall("GET", "/admin/v1/settings", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &SettingsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func SettingsCallerID(client_id string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("client_id", client_id)
	}
}

func SettingsFraudEmail(fraud_email string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("fraud_email", fraud_email)
	}
}

func SettingsFraudEnabled(fraud_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("fraud_enabled", Btoa(fraud_enabled))
	}
}

func SettingsInactiveUserExpiration(inactive_user_expiration int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("inactive_user_expiration", strconv.Itoa(inactive_user_expiration))
	}
}

func SettingsKeypressConfirm(keypress_confirm string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("keypress_confirm", keypress_confirm)
	}
}

func SettingsKeypressFraud(keypress_fraud string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("keypress_fraud", keypress_fraud)
	}
}

func SettingsLanguage(language string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("language", language)
	}
}

func SettingsLockoutExpireDuration(lockout_expire_duration string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("lockout_expire_duration", lockout_expire_duration)
	}
}

func SettingsLockoutThreshold(lockout_threshold int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("lockout_threshold", strconv.Itoa(lockout_threshold))
	}
}

func SettingsLogRetentionDays(log_retention_days int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("log_retention_days", strconv.Itoa(log_retention_days))
	}
}

func SettingsMinimumPasswordLength(minimum_password_length int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("minimum_password_length", strconv.Itoa(minimum_password_length))
	}
}

func SettingsMobileOTPEnabled(mobile_otp_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("mobile_otp_enabled", Btoa(mobile_otp_enabled))
	}
}

func SettingsName(name string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("name", name)
	}
}

func SettingsPasswordRequiredLowerAlpha(password_requires_lower_alpha bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("password_requires_lower_alpha", Btoa(password_requires_lower_alpha))
	}
}

func SettingsPasswordRequiredNumeric(password_requires_numeric bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("password_requires_numeric", Btoa(password_requires_numeric))
	}
}

func SettingsPasswordRequiresSpecial(password_requires_special bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("password_requires_special", Btoa(password_requires_special))
	}
}

func SettingsPasswordRequiresUpperAlpha(password_requires_upper_alpha bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("password_requires_upper_alpha", Btoa(password_requires_upper_alpha))
	}
}

func SettingsPushEnabled(push_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("push_enabled", Btoa(push_enabled))
	}
}

func SettingsSMSBatch(sms_batch int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("sms_batch", strconv.Itoa(sms_batch))
	}
}

func SettingsSMSEnabled(sms_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("sms_enabled", Btoa(sms_enabled))
	}
}

func SettingsSMSExpiration(sms_expiration int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("sms_expiration", strconv.Itoa(sms_expiration))
	}
}

func SettingsSMSMessage(sms_message string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("sms_message", sms_message)
	}
}

func SettingsSMSRefresh(sms_refresh int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("sms_refresh", strconv.Itoa(sms_refresh))
	}
}

func SettingsTelephonyWarningMin(telephony_warning_min int) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("telephony_warning_min", strconv.Itoa(telephony_warning_min))
	}
}

func SettingsTimezone(timezone string) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("timezone", timezone)
	}
}

func SettingsUserTelephonyCostMax(user_telephony_cost_max float32) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("user_telephony_cost_max", fmt.Sprintf("%.1f", user_telephony_cost_max))
	}
}

func SettingsVoiceEnabled(voice_enabled bool) func(*url.Values) {
	return func(opts *url.Values) {
		opts.Set("voice_enabled", Btoa(voice_enabled))
	}
}

//ModifySettings
//Required parameters: none
//Optional parameters: SettingsCallerID, SettingsFraudEmail, SettingsFraudEnabled, SettingsInactiveUserExpiration, SettingsKeypressConfirm, SettingsKeypressFraud, SettingsLanguage, SettingsLockoutThreshold, SettingsLogRetentionDays, SettingsMinimumPasswordLength, SettingsMobileOTPEnabled, SettingsName, SettingsPasswordRequiredLowerAlpha, SettingsPasswordRequiredNumeric, SettingsPasswordRequiresSpecial, SettingsPasswordRequiresUpperAlpha, SettingsPushEnabled, SettingsSMSBatch, SettingsSMSEnabled, SettingsSMSExpiration, SettingsSMSMessage, SettingsSMSRefresh, SettingsTelephonyWarningMin, SettingsTimezone, SettingsUserTelephonyCostMax, SettingsVoiceEnabled
func (api *AdminApi) ModifySettings(options ...func(*url.Values)) (*SettingsResult, error) {
	opts := url.Values{}
	for _, o := range options {
		o(&opts)
	}
	_, body, err := api.SignedCall("POST", "/admin/v1/settings", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &SettingsResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type MobileLogoResult struct {
	StatResult
	png *[]byte
}

//Retreive Duo Mobile Logo
//Required parameters: none
//Optional parameters: none
func (api *AdminApi) RetreiveDuoMobileLogo() (*MobileLogoResult, error) {
	resp, body, err := api.SignedCall("GET", "/admin/v1/logo", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		ret := &MobileLogoResult{StatResult: StatResult{Stat: "OK"},
			png: &body}
		return ret, nil
	}
	ret := &MobileLogoResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Modify Duo Mobile Logo
//Required parameters: logo
//Optional parameters: none
func (api *AdminApi) ModifyDuoMobileLogo(logo []byte) (*StatResult, error) {
	opts := url.Values{}
	opts.Set("logo", b64.StdEncoding.EncodeToString(logo))
	_, body, err := api.SignedCall("POST", "/admin/v1/logo", opts, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &StatResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

//Delete Duo Mobile Logo
//Required parameters: none
//Optional parameters: none
func (api *AdminApi) DeleteDuoMobileLogo() (*StatResult, error) {
	_, body, err := api.SignedCall("DELETE", "/admin/v1/logo", nil, duoapi.UseTimeout)
	if err != nil {
		return nil, err
	}
	ret := &StatResult{}
	if err = json.Unmarshal(body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
