package servicenow

import (
	"strconv"
)

type NowContact struct {
	*ServiceNow
}

func newNowContact(sn *ServiceNow) *NowContact {
	return &NowContact{sn}
}

type ContactGetResponse struct {
	Result Contact `json:"result"`
}
type ContactListResponse struct {
	Result []Contact `json:"result"`
}
type Contact struct {
	Account                 string `json:"account"`
	Active                  string `json:"active"`
	AgentStatus             string `json:"agent_status"`
	Building                string `json:"building"`
	CalendarIntegration     string `json:"calendar_integration"`
	City                    string `json:"city"`
	Company                 string `json:"company"`
	CostCenter              string `json:"cost_center"`
	Country                 string `json:"country"`
	DateFormat              string `json:"date_format"`
	DefaultPerspective      string `json:"default_perspective"`
	Department              string `json:"department"`
	EduStatus               string `json:"edu_status"`
	Email                   string `json:"email"`
	EmployeeNumber          string `json:"employee_number"`
	EnableMultifactorAuthn  string `json:"enable_multifactor_authn"`
	FailedAttempts          string `json:"failed_attempts"`
	FirstName               string `json:"first_name"`
	Gender                  string `json:"gender"`
	GeolocationTracked      string `json:"geolocation_tracked"`
	HomePhone               string `json:"home_phone"`
	InternalIntegrationUser string `json:"internal_integration_user"`
	Introduction            string `json:"introduction"`
	LastLogin               string `json:"last_login"`
	LastLoginDevice         string `json:"last_login_device"`
	LastLoginTime           string `json:"last_login_time"`
	LastName                string `json:"last_name"`
	LastPositionUpdate      string `json:"last_position_update"`
	Latitude                string `json:"latitude"`
	LdapServer              string `json:"ldap_server"`
	Location                string `json:"location"`
	LockedOut               string `json:"locked_out"`
	Longitude               string `json:"longitude"`
	Manager                 string `json:"manager"`
	MiddleName              string `json:"middle_name"`
	MobilePhone             string `json:"mobile_phone"`
	Name                    string `json:"name"`
	Notification            string `json:"notification"`
	OnSchedule              string `json:"on_schedule"`
	Phone                   string `json:"phone"`
	Photo                   string `json:"photo"`
	PreferredLanguage       string `json:"preferred_language"`
	Roles                   string `json:"roles"`
	Schedule                string `json:"schedule"`
	Source                  string `json:"source"`
	State                   string `json:"state"`
	Street                  string `json:"street"`
	SysClassName            string `json:"sys_class_name"`
	SysCreatedBy            string `json:"sys_created_by"`
	SysCreatedOn            string `json:"sys_created_on"`
	SysDomain               string `json:"sys_domain"`
	SysDomainPath           string `json:"sys_domain_path"`
	SysID                   string `json:"sys_id"`
	SysModCount             string `json:"sys_mod_count"`
	SysTags                 string `json:"sys_tags"`
	SysUpdatedBy            string `json:"sys_updated_by"`
	SysUpdatedOn            string `json:"sys_updated_on"`
	TimeFormat              string `json:"time_format"`
	TimeSheetPolicy         string `json:"time_sheet_policy"`
	TimeZone                string `json:"time_zone"`
	Title                   string `json:"title"`
	UserName                string `json:"user_name"`
	Vip                     string `json:"vip"`
	WebServiceAccessOnly    string `json:"web_service_access_only"`
	Zip                     string `json:"zip"`
}

// List all contacts.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/contact-api.html#title_contact-GET
func (sn *NowContact) List(limit, offset int) (*ContactListResponse, error) {
	var result ContactListResponse
	endpointUrl := sn.baseURL.JoinPath("api/now/contact")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	endpointUrl.RawQuery = queryUrl.Encode()

	err := sn.doAPI("GET", endpointUrl.String(), &result)
	return &result, err
}

// Get contact details.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/contact-api.html#title_contact-GET-id
func (sn *NowContact) Get(sysId string) (*ContactGetResponse, error) {
	var result ContactGetResponse
	endpointUrl := sn.baseURL.JoinPath("api/now/contact").JoinPath(sysId)
	err := sn.doAPI("GET", endpointUrl.String(), &result)
	return &result, err
}
