package private

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dwarvesf/go-threads/model"
)

func (t PrivateAPI) UnFollowUser(id int) (*model.FollowUserResponse, error) {
	url := fmt.Sprintf("%s/friendships/destroy/%d/", InstagramAPI, id)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req = updateRequestHeader(t.APIToken, req)

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

	var result model.FollowUserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
