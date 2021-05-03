package todoist

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const DefautlRestUrl string = "https://api.todoist.com/rest/v1"
const DefaultSyncurl string = "https://todoist.com/api/v8"

type Client struct {
	HTTPClient *http.Client
	Host       string
	ApiKey     string
	Base       string
	Req        *http.Request
}

func NewClient(apiKey *string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: time.Second * 10},
		ApiKey:     *apiKey,
	}
}

func (c *Client) newRequest(path, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", DefautlRestUrl, path), body)
	if err != nil {
		return nil, err
	}
	return req, err
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
