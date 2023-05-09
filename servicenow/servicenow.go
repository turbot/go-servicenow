package servicenow

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type ServiceNow struct {
	baseURL              *url.URL
	bearerToken          string
	basicAuth            string
	NowContact           *NowContact
	NowConsumer          *NowConsumer
	NowTable             *NowTable
	SnChgRestChange      *SnChgRestChange
	SnChgRestChangeModel *SnChgRestChangeModel
	SnChgRestChangeTask  *SnChgRestChangeTask
	SnChgRestChangeCi    *SnChgRestChangeCi
}

type Config struct {
	InstanceURL  string
	BasicAuth    bool
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

	sn := &ServiceNow{
		baseURL: baseURL,
	}

	if config.BasicAuth {
		sn.basicAuth = base64.StdEncoding.EncodeToString([]byte(config.Username + ":" + config.Password))
	} else {
		resp := &OAuthTokenResponse{}
		err = authenticate(&config, resp)
		if err != nil {
			return nil, err
		}
		sn.bearerToken = resp.AccessToken
	}

	sn.NowContact = newNowContact(sn)
	sn.NowConsumer = newNowConsumer(sn)
	sn.NowTable = newNowTable(sn)
	sn.SnChgRestChange = newSnChgRestChange(sn)
	sn.SnChgRestChangeModel = newSnChgRestChangeModel(sn)
	sn.SnChgRestChangeTask = newSnChgRestChangeTask(sn)
	sn.SnChgRestChangeCi = newSnChgRestChangeCi(sn)

	return sn, nil
}

// Authenticates the client and returns an API token to be used on API calls.
func authenticate(config *Config, resp interface{}) error {
	endpointUrl, _ := url.Parse(config.InstanceURL)
	endpointUrl = endpointUrl.JoinPath("oauth_token.do")

	payloadParameters := url.Values{
		"grant_type":    {config.GrantType},
		"client_id":     {config.ClientID},
		"client_secret": {config.ClientSecret},
		"username":      {config.Username},
		"password":      {config.Password},
	}
	payload := strings.NewReader(payloadParameters.Encode())

	header := http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	return DoRequest("POST", endpointUrl.String(), header, payload, &resp)
}

// Execute API calls
func (sn *ServiceNow) doAPI(method string, endpointUrl string, result interface{}) error {
	header := make(map[string][]string)
	if sn.basicAuth != "" {
		header["Authorization"] = []string{fmt.Sprintf("Basic %s", sn.basicAuth)}
	}
	if sn.bearerToken != "" {
		header["Authorization"] = []string{fmt.Sprintf("Bearer %s", sn.bearerToken)}
	}
	return DoRequest(method, endpointUrl, header, nil, &result)
}

func DoRequest(method string, endpointUrl string, header http.Header, payload io.Reader, result interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest(method, endpointUrl, payload)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}

	req.Header = header
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		httpError := struct {
			Error struct {
				Message string `json:"message"`
				Detail  string `json:"detail"`
			} `json:"error"`
		}{}
		err := json.Unmarshal(body, &httpError)

		if err != nil {
			httpError := struct {
				Error            string `json:"error"`
				ErrorDescription string `json:"error_description"`
			}{}
			err := json.Unmarshal(body, &httpError)
			if err != nil {
				return &HTTPError{Code: res.StatusCode, Message: string(body)}
			}
			return &HTTPError{Code: res.StatusCode, Message: httpError.ErrorDescription}
		}
		return &HTTPError{Code: res.StatusCode, Message: httpError.Error.Message, Detail: httpError.Error.Detail}
	}

	if isInstanceHibernating(body) {
		return fmt.Errorf("ServiceNow instance is hibernating")
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response payload: %w", err)
	}
	return nil
}

type HTTPError struct {
	Code    int
	Message string
	Detail  string
}

func (e *HTTPError) Error() string {
	if e.Code == 0 {
		e.Code = http.StatusInternalServerError
	}
	if e.Message != "" {
		return fmt.Sprintf("Got %d (%s) %s", e.Code, e.Message, e.Detail)
	}
	return fmt.Sprintf("Got %d (%s)", e.Code, http.StatusText(e.Code))
}

func isInstanceHibernating(body []byte) bool {
	re := regexp.MustCompile("Your instance is hibernating")
	return re.Match(body)
}
