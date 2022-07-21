package crawler

import (
	"io"
	"net/http"
	"time"
)

func GetHttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func GetResponseBody(endPoint string, client *http.Client) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response.Body, nil
}
