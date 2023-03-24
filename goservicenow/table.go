package goservicenow

import (
	"fmt"
	"net/http"
	"strconv"
)

func (c *Client) ListTable(tableName string, limit int, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s", tableName))

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	endpointUrl.RawQuery = queryUrl.Encode()

	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}

func (c *Client) GetTable(tableName string, sysId string, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s/%s", tableName, sysId))
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}
