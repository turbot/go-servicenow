package servicenow

import (
	"fmt"
	"strconv"
)

type NowTable struct {
	*ServiceNow
}

func newNowTable(sn *ServiceNow) *NowTable {
	return &NowTable{sn}
}

// List all tables.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/c_TableAPI.html#title_table-GET
func (sn *NowTable) List(tableName string, limit int, offset int, query string, excludeReferenceLink bool, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s", tableName))

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	queryUrl.Add("sysparm_exclude_reference_link", strconv.FormatBool(excludeReferenceLink))
	if query != "" {
		queryUrl.Add("sysparm_query", query)
	}
	endpointUrl.RawQuery = queryUrl.Encode()

	return sn.doAPI("GET", endpointUrl.String(), &result)
}

// Get table details.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/c_TableAPI.html#title_table-GET-id
func (sn *NowTable) Get(tableName string, sysId string, excludeReferenceLink bool, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s/%s", tableName, sysId))
	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_exclude_reference_link", strconv.FormatBool(excludeReferenceLink))
	endpointUrl.RawQuery = queryUrl.Encode()
	return sn.doAPI("GET", endpointUrl.String(), &result)
}
