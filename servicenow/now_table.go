package servicenow

import (
	"fmt"
	"net/http"
	"strconv"
)

type NowTable struct {
	*ServiceNow
}

func newNowTable(sn *ServiceNow) *NowTable {
	return &NowTable{sn}
}

// List all consumers.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/c_TableAPI.html#title_table-GET
func (sn *NowTable) List(tableName string, limit int, offset int, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s", tableName))

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	endpointUrl.RawQuery = queryUrl.Encode()

	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}
	return sn.doAPI(*req, result)
}

// Read consumer details.
//
// See: https://docs.servicenow.com/bundle/tokyo-application-development/page/integrate/inbound-rest/concept/c_TableAPI.html#title_table-GET-id
func (sn *NowTable) Read(tableName string, sysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s/%s", tableName, sysId))
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}
	return sn.doAPI(*req, result)
}
