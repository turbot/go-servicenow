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

func (sn *NowTable) List(tableName string, limit int, result interface{}) error {
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

func (sn *NowTable) Read(tableName string, sysId string, result interface{}) error {
	endpointUrl := sn.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s/%s", tableName, sysId))
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return sn.doAPI(*req, result)
}
