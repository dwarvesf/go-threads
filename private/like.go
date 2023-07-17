package private

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (t PrivateAPI) LikeThread(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/media/%s_%d/like/", InstagramAPI, id, t.UserID)
	req, err := http.NewRequest("POST", url, nil)
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

	var likingInfo map[string]interface{}
	err = json.Unmarshal(respBody, &likingInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return likingInfo, nil
}

func (t PrivateAPI) UnLikeThread(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/media/%s_%d/unlike/", InstagramAPI, id, t.UserID)
	req, err := http.NewRequest("POST", url, nil)
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

	var likingInfo map[string]interface{}
	err = json.Unmarshal(respBody, &likingInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return likingInfo, nil
}
