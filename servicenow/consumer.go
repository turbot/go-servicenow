package servicenow

import (
	"fmt"
	"net/http"
	"strconv"
)

type ConsumerGetResponse struct {
	Result Consumer `json:"result"`
}
type ConsumerListResponse struct {
	Result []Consumer `json:"result"`
}
type Consumer struct {
	Country           string `json:"country"`
	Notes             string `json:"notes"`
	Gender            string `json:"gender"`
	City              string `json:"city"`
	Prefix            string `json:"prefix"`
	SysUpdatedOn      string `json:"sys_updated_on"`
	Suffix            string `json:"suffix"`
	Title             string `json:"title"`
	Number            string `json:"number"`
	Notification      string `json:"notification"`
	SysID             string `json:"sys_id"`
	BusinessPhone     string `json:"business_phone"`
	SysUpdatedBy      string `json:"sys_updated_by"`
	MobilePhone       string `json:"mobile_phone"`
	Street            string `json:"street"`
	SysCreatedOn      string `json:"sys_created_on"`
	SysDomain         string `json:"sys_domain"`
	State             string `json:"state"`
	Fax               string `json:"fax"`
	FirstName         string `json:"first_name"`
	Email             string `json:"email"`
	PreferredLanguage string `json:"preferred_language"`
	SysCreatedBy      string `json:"sys_created_by"`
	Zip               string `json:"zip"`
	HomePhone         string `json:"home_phone"`
	TimeFormat        string `json:"time_format"`
	SysModCount       string `json:"sys_mod_count"`
	LastName          string `json:"last_name"`
	Photo             string `json:"photo"`
	Active            string `json:"active"`
	MiddleName        string `json:"middle_name"`
	TimeZone          string `json:"time_zone"`
	SysTags           string `json:"sys_tags"`
	Name              string `json:"name"`
	Household         string `json:"household"`
	DateFormat        string `json:"date_format"`
	User              string `json:"user"`
	Primary           string `json:"primary"`
}

func (c *Client) GetConsumers(limit, offset int) (*ConsumerListResponse, error) {
	var result ConsumerListResponse
	err := c.retrieveConsumers(limit, offset, &result)
	return &result, err
}

func (c *Client) GetConsumer(sysId string) (*ConsumerGetResponse, error) {
	var result ConsumerGetResponse
	err := c.retrieveConsumer(sysId, &result)
	return &result, err
}

func (c *Client) retrieveConsumers(limit, offset int, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath("api/now/consumer")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	queryUrl.Add("sysparm_offset", strconv.Itoa(offset))
	endpointUrl.RawQuery = queryUrl.Encode()

	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}

func (c *Client) retrieveConsumer(sysId string, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath("api/now/consumer").JoinPath(sysId)
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}
