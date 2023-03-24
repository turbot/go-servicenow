package goservicenow

import (
	"fmt"
	"net/http"
	"strconv"
)

type ArticleGetResponse struct {
	Result Article `json:"result"`
}
type ArticleListResponse struct {
	Result ArticleResult `json:"result"`
}
type Status struct {
	Code float64 `json:"code"`
}
type Meta struct {
	Start     float64 `json:"start"`
	End       float64 `json:"end"`
	Fields    string  `json:"fields"`
	Query     string  `json:"query"`
	Filter    string  `json:"filter"`
	Kb        string  `json:"kb"`
	Language  string  `json:"language"`
	Count     float64 `json:"count"`
	TsQueryID string  `json:"ts_query_id"`
	Status    Status  `json:"status"`
}
type ShortDescription struct {
	DisplayValue string `json:"display_value"`
	Name         string `json:"name"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	Value        string `json:"value"`
}
type SysClassName struct {
	DisplayValue string `json:"display_value"`
	Name         string `json:"name"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	Value        string `json:"value"`
}
type Fields struct {
	ShortDescription ShortDescription `json:"short_description"`
	SysClassName     SysClassName     `json:"sys_class_name"`
}
type Article struct {
	// List returned fields
	Link    string  `json:"link"`
	Rank    int     `json:"rank"`
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Snippet string  `json:"snippet"`
	Score   float64 `json:"score"`
	Number  string  `json:"number"`
	Fields  Fields  `json:"fields"`

	// Get returned fields
	Content            string        `json:"content"`
	Template           bool          `json:"template"`
	SysID              string        `json:"sys_id"`
	ShortDescription   string        `json:"short_description"`
	DisplayAttachments bool          `json:"display_attachments"`
	EmbeddedContent    []interface{} `json:"embedded_content"`
}
type ArticleResult struct {
	Meta     Meta      `json:"meta"`
	Articles []Article `json:"articles"`
}

func (c *Client) GetArticles(limit, offset int) (*ArticleListResponse, error) {
	var result ArticleListResponse
	err := c.retrieveArticles(limit, offset, &result)
	return &result, err
}

func (c *Client) GetArticle(sysId string) (*ArticleGetResponse, error) {
	var result ArticleGetResponse
	err := c.retrieveArticle(sysId, &result)
	return &result, err
}

func (c *Client) retrieveArticles(limit, offset int, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath("api/sn_km_api/knowledge/articles")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("limit", strconv.Itoa(limit))
	queryUrl.Add("offset", strconv.Itoa(offset))
	endpointUrl.RawQuery = queryUrl.Encode()

	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}

func (c *Client) retrieveArticle(sysId string, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath("api/sn_km_api/knowledge/articles").JoinPath(sysId)
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}