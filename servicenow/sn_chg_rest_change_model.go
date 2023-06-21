package servicenow

import (
	"fmt"
	"strconv"
)

type SnChgRestChangeModel struct {
	*ServiceNow
}

func newSnChgRestChangeModel(sn *ServiceNow) *SnChgRestChangeModel {
	return &SnChgRestChangeModel{sn}
}

// List all change models.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-model
func (sn *SnChgRestChangeModel) List(limit int, offset int, query string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath("api/sn_chg_rest/change/model")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	if query != "" {
		queryUrl.Add("sysparm_query", query)
	}
	endpointUrl.RawQuery = queryUrl.Encode()

	return sn.doAPI("GET", endpointUrl.String(), &result)
}

// Get change model details.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-model-sys_id
func (sn *SnChgRestChangeModel) Get(sysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/sn_chg_rest/change/model/%s", sysId))
	return sn.doAPI("GET", endpointUrl.String(), &result)
}
