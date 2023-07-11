package private

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dwarvesf/go-threads/model"
)

func (t PrivateAPI) GetUserFollowers(id int) (*model.UserFollowersResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/friendships/%d/followers/", InstagramAPI, id), nil)
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d, response: %s", resp.StatusCode, string(respBody))
	}

	var followers model.UserFollowersResponse
	err = json.Unmarshal(respBody, &followers)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &followers, nil
}

func (t PrivateAPI) GetFollowers() (*model.UserFollowersResponse, error) {
	return t.GetUserFollowers(t.UserID)
}
