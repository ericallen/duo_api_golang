package adminapi

import "github.com/ericallen/duo_api_golang"

type AdminApi struct {
	duoapi.DuoApi
}

//Build a new Duo Admin API object.
// api is a duoapi.DuoApi object used to make the Duo Rest API calls.
// Example: adminapi.NewAdminApi(*duoapi.NewDuoApi(ikey,skey,host,userAgent,duoapi.SetTimeout(10*time.second)))
func NewAdminApi(api duoapi.DuoApi) *AdminApi {
	return &AdminApi{api}
}
