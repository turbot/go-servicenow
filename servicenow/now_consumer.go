package servicenow

import (
	"strconv"
)

type NowConsumer struct {
	*ServiceNow
}

func newNowConsumer(sn *ServiceNow) *NowConsumer {
	return &NowConsumer{sn}
}

type ConsumerGetResponse struct {
	Result Consumer `json:"result"`
}
type ConsumerListResponse struct {
	Result []Consumer `json:"result"`
}
type Consumer struct {
	Active            string `json:"active"`
	BusinessPhone     string `json:"business_phone"`
	City              string `json:"city"`
	Country           string `json:"country"`
	DateFormat        string `json:"date_format"`
	Email             string `json:"email"`
	Fax               string `json:"fax"`
	FirstName         string `json:"first_name"`
	Gender            string `json:"gender"`
	HomePhone         string `json:"home_phone"`
	Household         string `json:"household"`
	LastName          string `json:"last_name"`
	MiddleName        string `json:"middle_name"`
	MobilePhone       string `json:"mobile_phone"`
	Name              string `json:"name"`
	Notes             string `json:"notes"`
	Notification      string `json:"notification"`
	Number            string `json:"number"`
	Photo             string `json:"photo"`
	PreferredLanguage string `json:"preferred_language"`
	Prefix            string `json:"prefix"`
	Primary           string `json:"primary"`
	State             string `json:"state"`
	Street            string `json:"street"`
	Suffix            string `json:"suffix"`
	SysCreatedBy      string `json:"sys_created_by"`
	SysCreatedOn      string `json:"sys_created_on"`
	SysDomain         string `json:"sys_domain"`
	SysID             string `json:"sys_id"`
	SysModCount       string `json:"sys_mod_count"`
	SysTags           string `json:"sys_tags"`
	SysUpdatedBy      string `json:"sys_updated_by"`
	SysUpdatedOn      string `json:"sys_updated_on"`
	TimeFormat        string `json:"time_format"`
	TimeZone          string `json:"time_zone"`
	Title             string `json:"title"`
	User              string `json:"user"`
	Zip               string `json:"zip"`
}

// List all consumers.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/consumer-api.html#title_consumer-GET
func (sn *NowConsumer) List(limit, offset int) (*ConsumerListResponse, error) {
	var result ConsumerListResponse
	endpointUrl := sn.baseURL.JoinPath("api/now/consumer")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	endpointUrl.RawQuery = queryUrl.Encode()

	err := sn.doAPI("GET", endpointUrl.String(), &result)
	return &result, err
}

// Read consumer details.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/consumer-api.html#title_consumer-GET-id
func (sn *NowConsumer) Read(sysId string) (*ConsumerGetResponse, error) {
	var result ConsumerGetResponse
	endpointUrl := sn.baseURL.JoinPath("api/now/consumer").JoinPath(sysId)
	err := sn.doAPI("GET", endpointUrl.String(), &result)
	return &result, err
}
