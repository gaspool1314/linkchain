package daemon

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/lianxiangcloud/linkchain/libs/log"
	"github.com/lianxiangcloud/linkchain/wallet/config"
)

// type HttpMethod string

// const (
// 	GET  HttpMethod = "GET"
// 	POST HttpMethod = "POST"
// )

type DaemonClient struct {
	Addr string
	// Login string

	// Trusted bool
	// Testnet bool

	HttpClient *http.Client
}

var gDaemonClient *DaemonClient

const (
	defaultDialTimeout = 10 * time.Second
	keepAliveInterval  = 30 * time.Second
)

func InitDaemonClient(daemonConfig *config.DaemonConfig) {
	gDaemonClient = &DaemonClient{
		Addr: daemonConfig.PeerRPC,
		// Login:   daemonConfig.Login,
		// Trusted: daemonConfig.Trusted,
		// Testnet: daemonConfig.Testnet,
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
		ResponseHeaderTimeout: 2 * time.Minute,
		DisableCompression:    true,
		DisableKeepAlives:     false,
		IdleConnTimeout:       2 * time.Minute,
		MaxIdleConns:          4,
		MaxIdleConnsPerHost:   2,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			dialer := &net.Dialer{Timeout: defaultDialTimeout, KeepAlive: keepAliveInterval}
			return dialer.DialContext(ctx, network, addr)
		},
	}
	gDaemonClient.HttpClient = &http.Client{
		Transport: transport,
	}
}

// CallJSONRPC call  /json_rpc func
// curl -X POST http://127.0.0.1:18081/json_rpc -d '{"jsonrpc":"2.0","id":"0","method":"get_block","params":{"height":912345}}' -H 'Content-Type: application/json'
func CallJSONRPC(method string, params interface{}) ([]byte, error) {
	urlPath := ""
	if len(method) >= 4 {
		urlPath = method[4:]
	}

	url := fmt.Sprintf("%s/%s", gDaemonClient.Addr, urlPath)

	requestData := make(map[string]interface{})

	requestData["jsonrpc"] = "2.0"
	requestData["id"] = 1
	requestData["method"] = method
	requestData["params"] = params

	client := gDaemonClient.HttpClient
	data, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	log.Debug("CallJSONRPC", "url", url, "data", string(data))
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("NewRequest: err=%v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("nc", "IN")
	req = req.WithContext(context.Background())
	resp, err := client.Do(req)
	if err != nil {
		// log.Error("CallJSONRPC client.Do", "err", err)
		return nil, fmt.Errorf("client.Do: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, nil
}
