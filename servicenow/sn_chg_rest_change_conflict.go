package servicenow

import (
	"fmt"
)

type SnChgRestChangeConflict struct {
	*ServiceNow
}

func newSnChgRestChangeConflict(sn *ServiceNow) *SnChgRestChangeConflict {
	return &SnChgRestChangeConflict{sn}
}

// Get change conflict.
//
// See: https://docs.servicenow.com/bundle/rome-application-development/page/integrate/inbound-rest/concept/change-management-api.html#title_change-GET-conflict
func (sn *SnChgRestChangeConflict) Get(changeSysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/sn_chg_rest/change/%s/conflict", changeSysId))
	return sn.doAPI("GET", endpointUrl.String(), result)
}
