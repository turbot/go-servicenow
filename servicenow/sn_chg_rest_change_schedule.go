package servicenow

import (
	"fmt"
)

type SnChgRestChangeSchedule struct {
	*ServiceNow
}

func newSnChgRestChangeSchedule(sn *ServiceNow) *SnChgRestChangeSchedule {
	return &SnChgRestChangeSchedule{sn}
}

// Get change schedule.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-sched
func (sn *SnChgRestChangeSchedule) Get(changeSysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/sn_chg_rest/change/%s/schedule", changeSysId))
	return sn.doAPI("GET", endpointUrl.String(), result)
}
