package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mondaycodingclub/my-nginx/api/heartbeat"
	"io"
	"io/ioutil"
	"net/http"
)

type Interface interface {
	SyncHeartbeat(request *heartbeat.Request) (*heartbeat.Response, error)
}

type MasterClient struct {
	masterHost  string
	masterPort  int
	masterToken string
}

func NewMasterClient(masterHost string, masterPort int, masterToken string) *MasterClient {
	return &MasterClient{
		masterHost:  masterHost,
		masterPort:  masterPort,
		masterToken: masterToken,
	}
}

func (mc *MasterClient) SyncHeartbeat(request *heartbeat.Request) (*heartbeat.Response, error) {
	url := fmt.Sprintf("https://%s:%d/v1/agents/agent/heartbeats", mc.masterHost, mc.masterPort)
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(b)
	resp, err := http.Post(url, "", body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &heartbeat.Response{}
	if err = json.Unmarshal(content, response); err != nil {
		return nil, err
	}
	return response, nil
}
