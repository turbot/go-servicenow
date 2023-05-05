package servicenow

import (
	"fmt"
	"strconv"
)

type SnChgRestChangeCi struct {
	*ServiceNow
}

func newSnChgRestChangeCi(sn *ServiceNow) *SnChgRestChangeCi {
	return &SnChgRestChangeCi{sn}
}

// List change ci.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-ci
func (sn *SnChgRestChangeCi) List(changeSysId string, associationType string, limit int, offset int, query string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/sn_chg_rest/change/%s/ci", changeSysId))

	queryUrl := endpointUrl.Query()
	queryUrl.Add("association_type", associationType)
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	if query != "" {
		queryUrl.Add("sysparm_query", query)
	}
	endpointUrl.RawQuery = queryUrl.Encode()

	return sn.doAPI("GET", endpointUrl.String(), result)
}
