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
	Country           string `json:"country"`
	Notes             string `json:"notes"`
	Gender            string `json:"gender"`
	City              string `json:"city"`
	Prefix            string `json:"prefix"`
	SysUpdatedOn      string `json:"sys_updated_on"`
	Suffix            string `json:"suffix"`
	Title             string `json:"title"`
	Number            string `json:"number"`
	Notification      string `json:"notification"`
	SysID             string `json:"sys_id"`
	BusinessPhone     string `json:"business_phone"`
	SysUpdatedBy      string `json:"sys_updated_by"`
	MobilePhone       string `json:"mobile_phone"`
	Street            string `json:"street"`
	SysCreatedOn      string `json:"sys_created_on"`
	SysDomain         string `json:"sys_domain"`
	State             string `json:"state"`
	Fax               string `json:"fax"`
	FirstName         string `json:"first_name"`
	Email             string `json:"email"`
	PreferredLanguage string `json:"preferred_language"`
	SysCreatedBy      string `json:"sys_created_by"`
	Zip               string `json:"zip"`
	HomePhone         string `json:"home_phone"`
	TimeFormat        string `json:"time_format"`
	SysModCount       string `json:"sys_mod_count"`
	LastName          string `json:"last_name"`
	Photo             string `json:"photo"`
	Active            string `json:"active"`
	MiddleName        string `json:"middle_name"`
	TimeZone          string `json:"time_zone"`
	SysTags           string `json:"sys_tags"`
	Name              string `json:"name"`
	Household         string `json:"household"`
	DateFormat        string `json:"date_format"`
	User              string `json:"user"`
	Primary           string `json:"primary"`
}

// List all consumers.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/consumer-api.html#title_consumer-GET
func (sn *NowConsumer) List(limit, offset int) (*ConsumerListResponse, error) {
	var result ConsumerListResponse
	err := sn.retrieveConsumers(limit, offset, &result)
	return &result, err
}

// Read consumer details.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/consumer-api.html#title_consumer-GET-id
func (sn *NowConsumer) Read(sysId string) (*ConsumerGetResponse, error) {
	var result ConsumerGetResponse
	err := sn.retrieveConsumer(sysId, &result)
	return &result, err
}

func (sn *NowConsumer) retrieveConsumers(limit, offset int, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath("api/now/consumer")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	endpointUrl.RawQuery = queryUrl.Encode()

	return sn.doAPI("GET", endpointUrl.String(), result)
}

func (sn *NowConsumer) retrieveConsumer(sysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath("api/now/consumer").JoinPath(sysId)
	return sn.doAPI("GET", endpointUrl.String(), result)
}
