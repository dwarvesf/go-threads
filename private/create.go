package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/dwarvesf/go-threads/model"
)

func (t PrivateAPI) CreatePost(content model.CreatePostRequest) (*model.CreatePostResponse, error) {
	currentTimestamp := time.Now().Unix()
	userID := t.UserID

	parameters := map[string]interface{}{
		"publish_mode":       "text_post",
		"text_post_app_info": `{"reply_control":0}`,
		"timezone_offset":    t.TimezoneOffset,
		"source_type":        "4",
		"caption":            content.Caption,
		"_uid":               userID,
		"device_id":          t.AndroidDeviceID,
		"upload_id":          currentTimestamp,
		"device": map[string]interface{}{
			"manufacturer":    "OnePlus",
			"model":           "ONEPLUS+A3010",
			"android_version": 25,
			"android_release": "7.1.1",
		},
	}

	parametersJSON, err := json.Marshal(parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal parameters to JSON: %v", err)
	}

	encodedParameters := url.QueryEscape(string(parametersJSON))

	reqBody := fmt.Sprintf("signed_body=SIGNATURE.%s", encodedParameters)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/media/configure_text_only_post/", InstagramAPI), bytes.NewBufferString(reqBody))
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

	var createThreadResp model.CreatePostResponse
	err = json.Unmarshal(respBody, &createThreadResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return &createThreadResp, nil
}
