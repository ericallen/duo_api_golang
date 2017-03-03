package adminapi

import (
  "fmt"
  "github.com/ericallen/duo_api_golang"
)


type AdminsResult struct {
  Stat string
  Response []Admin
}

type AdminResult struct {
  Stat string
  Response Admin
}

type Admin struct {
  Name string
  Admin_ID string
  Email string
  Phone string
  Last_Login int64
  Role string
}

//Retreive Administrators
//Required parameters: none
//Optional parameters: none
func (api *AdminApi)RetreiveAdministrators()(*AdminsResult, error)  {
    _, body, err := api.SignedCall("GET", "/admin/v1/admins", nil, duoapi.UseTimeout)
    if err != nil {
      return nil, err
    }
    ret := &AdminsResult{}
    err = json.Unmarshal(body, ret); err != nil {
      return nil, err
    }
  return ret, nil
}

//Create Administrator
//Required parameters: email, password, name, phone
//Optional parameters: Role
func (api *AdminApi)CreateAdministrator(email, password, name, phone string, options ...func(*url.Values))(*AdminResult,error)  {
  opts := url.Values{}
  opts.Set("email", email)
  opts.Set("password", password)
  opts.Set("name", name)
  opts.Set("phone", phone)
  for _, o := range options {
    o(&opts)
  }
  _, body, err := api.SignedCall("POST", "/admin/v1/admins", opts, duoapi.UseTimeout)
  if err != nil {
    return nil, err
  }
  ret := &AdminResult{}
  if err = json.Unmarshal(body, ret); err != nil {
    return nil, err
  }
  return ret, nil
}

//Retreive Administator by ID
//Required parameters: administrator_id
//Optional parameters: none
func (api *AdminApi)RetreiveAdminstatorbyID(administrator_id string) (*AdminResult,error)  {
  path := fmt.Sprintf("/admin/v1/admins/%s", administrator_id)
  _, body, err := api.SignedCall("GET", path, nil, duoapi.UseTimeout)
  if err != nil {
    return nil, err
  }
  ret := &AdminResult{}
  if err = json.Unmarshal(body, ret); err != nil {
    return nil, err
  }
  return ret, nil
}

func AdminName(name string) func(*url.Values)  {
  return func(opts *url.Values) {
		opts.Set("name", name)
	}
}

func AdminPhone(phone string) func(*url.Values){
  return func(opts *url.Values){
    opts.Set("phone", phone)
  }
}

func AdminRole(role string) func(*url.Values){
  return func(opts *url.Values){
    opts.Set("role", role)
  }
}

//Modify administrator
//Required parameters: administrator_id
//Optional parameters: AdminName, AdminPhone, AdminPassword, AdminRole
func (api *AdminApi)ModifyAdministrator(administrator_id string, options ...func(*url.Values)) (*AdminResult, error) {
  opts := url.Values{}
  for _, o := range options {
    o(&opts)
  }
  path := fmt.Sprintf("/admin/v1/admins/%s", administrator_id)
  _, body, err := api.SignedCall("POST", path, opts, duoapi.UseTimeout)
  if err != nil {
    return nil, err
  }
  ret := &AdminResult{}
  if err = json.Unmarshal(body, ret); err != nil {
    return nil, err
  }
  return ret, nil
}

//Delete administrator
//Required parameters: administrator_id
//Optional parameters: none
func (api *AdminApi)DeleteAdministrator(administrator_id string)(*SimpleResponse, error)  {
  path := fmt.Sprintf("/admin/v1/admins/%s", administrator_id)
  _, body, err := api.SignedCall("DELETE", path, nil, duoapi.UseTimeout)
  if err != nil {
    return nil, err
  }
  ret := &SimpleResponse{}
  if err = json.Unmarshal(body, ret); err != nil {
    return nil, err
  }
  return ret, nil
}

//Reset Administrator Authentication Attempts
//Required parameters: administrator_id
//Optional parameters: none
func (api *AdminApi)ResetAdministratorAuthenticationAttemps(administrator_id string) (*SimpleResponse, error)  {
  path := fmt.Sprintf("/admin/v1/admins/%s/reset", administrator_id)
  _, body, err := api.SignedCall("POST", path, nil, duoapi.UseTimeout)
  if err != nil {
    return nil, err
  }
  ret := &SimpleResponse{}
  if err = json.Unmarshal(body, ret); err != nil {
    return nil, err
  }
  return ret, nil
}

type AdminActivationLink struct {
  Code string
  Email string
  Email_Sent int
  Link string
  Message string
  Valid_Days int
}

type ActivationLinkResult struct {
  Stat string
  Response AdminActivationLink
}

//Create Administrator Activation Link
//Required parameters: email
//Optional parameters: AdminSendEmail, AdminValidDays
func (api *AdminApi)CreateAdministratorActivationLink(email string, options ...func(*url.Values)) (*ActivationLinkResult, error)  {
  opts := url.Values{}
  for _, o := range options {
    o(&opts)
  }
  _, body, err := api.SignedCall("POST", "/admin/v1/admins/activate", opts, duoapi.UseTimeout)
  if err != nil {
    return nil, err
  }
  ret := &ActivationLinkResult{}
  if err = json.Unmarshal(body, ret); err != nil {
    return nil, err
  }
  return ret, nil
}
