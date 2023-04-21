package servicenow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ServiceNow struct {
	baseURL                  *url.URL
	servicenowToken          string
	NowContact               *NowContact
	NowConsumer              *NowConsumer
	NowTable                 *NowTable
	SnKmApiKnowledgeArticles *SnKmApiKnowledgeArticles
	SnChgRestChange          *SnChgRestChange
}

type Config struct {
	InstanceURL  string
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

type OAuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// New creates a new ServiceNow client by authenticating using the supplied credentials.
func New(config Config) (serviceNow *ServiceNow, err error) {
	baseURL, _ := url.Parse(config.InstanceURL)

	resp := &OAuthTokenResponse{}
	err = authenticate(config, resp)
	if err != nil {
		return nil, err
	}

	sn := &ServiceNow{
		baseURL:         baseURL,
		servicenowToken: resp.AccessToken,
	}
	sn.NowContact = newNowContact(sn)
	sn.NowConsumer = newNowConsumer(sn)
	sn.NowTable = newNowTable(sn)
	sn.SnKmApiKnowledgeArticles = newSnKmApiKnowledgeArticles(sn)
	sn.SnChgRestChange = newSnChgRestChange(sn)

	return sn, nil
}

// Authenticates the client and returns an API token to be used on API calls.
func authenticate(config Config, resp interface{}) error {
	endpointUrl, _ := url.Parse(config.InstanceURL)
	endpointUrl = endpointUrl.JoinPath("oauth_token.do")
	method := "POST"

	payloadParameters := url.Values{
		"grant_type":    {config.GrantType},
		"client_id":     {config.ClientID},
		"client_secret": {config.ClientSecret},
		"username":      {config.Username},
		"password":      {config.Password},
	}
	payload := strings.NewReader(payloadParameters.Encode())

	client := &http.Client{}
	req, err := http.NewRequest(method, endpointUrl.String(), payload)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpResp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode >= 300 {
		httpError := struct {
			Message string `json:"message"`
		}{}
		err = json.NewDecoder(httpResp.Body).Decode(&httpError)
		if err != nil {
			return fmt.Errorf("failed to decode json error response payload: %w", err)
		}
		return &HTTPError{Code: httpResp.StatusCode, Message: httpError.Message}
	}

	if resp != nil {
		err = json.NewDecoder(httpResp.Body).Decode(resp)
		if err != nil {
			return fmt.Errorf("failed to decode json error response payload: %w", err)
		}
	}
	return nil
}

// Execute API calls
func (sn *ServiceNow) doAPI(method string, endpointUrl string, result interface{}) error {
	req, err := http.NewRequest(method, endpointUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}

	client := &http.Client{}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", sn.servicenowToken))
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	defer res.Body.Close()

	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response payload: %w", err)
	}
	return nil
}

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	if e.Code == 0 {
		e.Code = http.StatusInternalServerError
	}
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("response %d (%s)", e.Code, http.StatusText(e.Code))
}
