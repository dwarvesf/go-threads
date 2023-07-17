package public

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/dwarvesf/go-threads/model"
)

func (s PublicAPI) GetThreadLikers(ID int) (*model.LikersResponse, error) {

	data := url.Values{
		"lsd":       []string{s.APPToken},
		"variables": []string{fmt.Sprintf(`{"mediaID": %d}`, ID)},
		"doc_id":    []string{"9360915773983802"},
	}

	req, err := http.NewRequest("POST", ThreadsAPIURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req = prepareHeaders(req, s.APPToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result model.LikersResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
