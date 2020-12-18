package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type AzureDevOpsBuildRQ struct {
	Def          Definition `json:"definition"`
	Params       string     `json:"parameters"`
	SourceBranch string     `json"sourceBranch`
}

type Definition struct {
	ID string `json:"id"`
}

//TODO: 此功能目前開發給azuredevops build pipeline使用, payload spec 有特規
//Callback ...
func Callback(endpoint string, status bool, extraInfo string) error {
	var rq AzureDevOpsBuildRQ
	json.Unmarshal([]byte(extraInfo), &rq)
	if status {
		rq.Params = "{\"signal\": 1}"
	} else {
		rq.Params = "{\"signal\": 0}"
	}
	rq.SourceBranch = os.Getenv("AZP_MAPPING_BRANCH")
	jsonRQ, _ := json.Marshal(rq)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonRQ))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+os.Getenv("AZP_PAT"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
