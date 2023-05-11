package servicenow

import (
	"fmt"
	"strconv"
)

type SnChgRestChangeTask struct {
	*ServiceNow
}

func newSnChgRestChangeTask(sn *ServiceNow) *SnChgRestChangeTask {
	return &SnChgRestChangeTask{sn}
}

// List all change tasks.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-task
func (sn *SnChgRestChangeTask) List(changeSysId string, limit int, offset int, query string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/sn_chg_rest/change/%s/task", changeSysId))

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	if query != "" {
		queryUrl.Add("sysparm_query", query)
	}
	endpointUrl.RawQuery = queryUrl.Encode()

	return sn.doAPI("GET", endpointUrl.String(), &result)
}
