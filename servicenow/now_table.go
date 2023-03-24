package servicenow

import (
	"fmt"
	"net/http"
	"strconv"
)

func (sn *ServiceNow) ListTable(tableName string, limit int, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s", tableName))

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	endpointUrl.RawQuery = queryUrl.Encode()

	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return sn.doAPI(*req, result)
}

func (sn *ServiceNow) GetTable(tableName string, sysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s/%s", tableName, sysId))
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return sn.doAPI(*req, result)
}
