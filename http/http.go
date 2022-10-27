package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func HttpClientInitialize() (*HttpClient, error) {

	httpClient := &HttpClient{
		client: &http.Client{},
	}

	return httpClient, nil
}

func (httpClient *HttpClient) Get(url string, to int) (string, int, error) {
	client := httpClient.client
	client.Timeout = time.Second * time.Duration(to)

	res, err := client.Get(url)
	if err != nil {
		return err.Error(), 500, err
	}
	defer res.Body.Close()

	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error(), res.StatusCode, err
	}

	str := string(bodyData)

	return str, res.StatusCode, err
}

func (httpClient *HttpClient) GetRequest(url string, to int) (map[string]interface{}, error) {
	client := httpClient.client
	client.Timeout = time.Second * time.Duration(to)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("http response : %v", res)
	return ResponseHandler(res)
}

func ResponseHandler(resp *http.Response) (map[string]interface{}, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return parseBody(resp)
	default:
		err := fmt.Errorf("http response error")
		return nil, err
	}
}

// response 가 정상일 경우 map 형태로 받아오도록 구성하였습니다
func parseBody(response *http.Response) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}
