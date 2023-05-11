package servicenow

import (
	"fmt"
	"strconv"
)

type SnChgRestChange struct {
	*ServiceNow
}

func newSnChgRestChange(sn *ServiceNow) *SnChgRestChange {
	return &SnChgRestChange{sn}
}

// List all changes.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-change
func (sn *SnChgRestChange) List(limit int, offset int, query string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath("api/sn_chg_rest/change")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	if query != "" {
		queryUrl.Add("sysparm_query", query)
	}
	endpointUrl.RawQuery = queryUrl.Encode()

	return sn.doAPI("GET", endpointUrl.String(), &result)
}

// Read change details.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-change-sys_id
func (sn *SnChgRestChange) Read(sysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/sn_chg_rest/change/%s", sysId))
	return sn.doAPI("GET", endpointUrl.String(), &result)
}
