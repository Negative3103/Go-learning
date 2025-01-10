package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout must be greater than 0")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAssets(name string) (AssetData, error) {
	url := fmt.Sprintf("http://api.coincap.io/v2/assets/%s", name)
	response, err := c.client.Get(url)
	if err != nil {
		return AssetData{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return AssetData{}, err
	}

	var r assetResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return AssetData{}, err
	}
	return r.Data, nil
}
