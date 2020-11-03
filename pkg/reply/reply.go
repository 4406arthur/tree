package reply

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/pquerna/ffjson/ffjson"
)

// ControllerSDK used to reply job status
type ControllerSDK struct {
	httpClient *httpclient.Client
	url        string
}

type status struct {
	JobID   string `json:"jobID"`
	Success bool   `json:"success"`
}

// NewControllerSDK ...
func NewControllerSDK(httpCli *httpclient.Client, url string) *ControllerSDK {
	return &ControllerSDK{
		httpClient: httpCli,
		url:        url,
	}
}

// Reply ...
func (c *ControllerSDK) Reply(jobID string, success bool) error {
	endpoint := c.url + "/finish"
	payload, _ := ffjson.Marshal(&status{
		JobID:   jobID,
		Success: success,
	})
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Post(
		endpoint,
		bytes.NewBuffer(payload),
		headers,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return errors.New("wrong http status code")
	}

	return nil
}
