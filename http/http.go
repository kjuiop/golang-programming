package http

import (
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
	//client := &http.Client{}
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
