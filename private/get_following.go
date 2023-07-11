package private

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dwarvesf/go-threads/model"
)

func (t PrivateAPI) GetUserFollowing(id int) (*model.UserFollowingResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/friendships/%d/following/", InstagramAPI, id), nil)
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

	var following model.UserFollowingResponse
	err = json.Unmarshal(respBody, &following)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &following, nil
}

func (t PrivateAPI) GetFollowing() (*model.UserFollowingResponse, error) {
	return t.GetUserFollowing(t.UserID)
}
