package private

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dwarvesf/go-threads/model"
)

func (t PrivateAPI) SearchUser(query string) (*model.UserResponse, error) {
	url := fmt.Sprintf("%s/users/search/?q=%s", InstagramAPI, query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req = updateRequestHeader(t.APIToken, req)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var users model.UserResponse
	err = json.Unmarshal(respBody, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &users, nil
}
