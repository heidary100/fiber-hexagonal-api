package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearchResponse struct {
	Link string `json:"link"`
}

func Search(q string) ([]SearchResponse, error) {
	var record []SearchResponse
	safeName := url.QueryEscape(q)
	endPoint := fmt.Sprintf("http://194.5.206.87:5040/api/search/google?&q=%s", safeName)
	response, err := http.Get(endPoint)

	if err != nil {
		return record, err
	}

	if err = json.NewDecoder(response.Body).Decode(&record); err != nil {
		return record, err
	}
	return record, nil
}
