package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/dwarvesf/go-threads/model"
)

const InstagramAPI = "https://i.instagram.com/api/v1"

func (t *PrivateAPI) Auth() (*model.AuthResponse, error) {
	blockVersion := "5f56efad68e1edec7801f630b5c122704ec5378adbee6609a448f105f34a9c73"

	parameters := map[string]interface{}{
		"client_input_params": map[string]interface{}{
			"password":      t.Password,
			"contact_point": t.Username,
			"device_id":     t.AndroidDeviceID,
		},
		"server_params": map[string]interface{}{
			"credential_type": "password",
			"device_id":       t.AndroidDeviceID,
		},
	}

	parametersAsBytes, err := json.Marshal(parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal parameters to JSON: %v", err)
	}

	bkClientContext := map[string]interface{}{
		"bloks_version": blockVersion,
		"styles_id":     "instagram",
	}

	bkClientContextAsBytes, err := json.Marshal(bkClientContext)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal bkClientContext to JSON: %v", err)
	}

	params := url.QueryEscape(string(parametersAsBytes))
	bkClientContextStr := url.QueryEscape(string(bkClientContextAsBytes))
	reqBody := fmt.Sprintf("params=%s&bk_client_context=%s&bloks_versioning_id=%s", params, bkClientContextStr, blockVersion)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/bloks/apps/com.bloks.www.bloks.caa.login.async.send_login_request/", InstagramAPI), bytes.NewBufferString(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Barcelona 289.0.0.77.109 Android")
	req.Header.Set("Sec-Fetch-Site", "same-origin")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d, response: %s", resp.StatusCode, string(respBody))
	}

	raw := string(respBody)
	fmt.Println("raw")
	fmt.Println(raw)
	bearerKeyPosition := strings.Index(raw, "Bearer IGT:2:")
	key := raw[bearerKeyPosition:]
	backslashKeyPosition := strings.Index(key, "\\\\")

	token := key[13:backslashKeyPosition]
	t.APIToken = token
	return &model.AuthResponse{Token: token}, nil
}
