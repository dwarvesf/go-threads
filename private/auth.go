package private

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/dwarvesf/go-threads/model"
	"github.com/dwarvesf/go-threads/util"
	"github.com/google/uuid"
)

type publicKeyResponse struct {
	KeyID int
	Key   []byte
}

func getInstagramPublicKey() (*publicKeyResponse, error) {
	parameters := struct {
		ID string `json:"id"`
	}{
		ID: uuid.NewString(),
	}

	parametersAsBytes, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}

	encodedParameters := url.QueryEscape(string(parametersAsBytes))
	req, err := http.NewRequest("POST", InstagramAPI+"/qe/sync/", strings.NewReader(fmt.Sprintf("params=%s", encodedParameters)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Barcelona 289.0.0.77.109 Android")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	publicKeyKeyID := resp.Header.Get("ig-set-password-encryption-key-id")
	publicKey := resp.Header.Get("ig-set-password-encryption-pub-key")

	publicKeyKeyIDAsInt, err := strconv.Atoi(publicKeyKeyID)
	if err != nil {
		return nil, err
	}
	rawDecodedText, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	return &publicKeyResponse{publicKeyKeyIDAsInt, rawDecodedText}, nil
}

func Auth(username string, password string, deviceID string) (*model.AuthResponse, error) {
	pub, err := getInstagramPublicKey()
	if err != nil {
		return nil, err
	}
	encryptedPassword, timestamStr, _ := util.EncryptPassword(password, pub.KeyID, pub.Key)
	blockVersion := "5f56efad68e1edec7801f630b5c122704ec5378adbee6609a448f105f34a9c73"

	parameters := map[string]interface{}{
		"client_input_params": map[string]interface{}{
			"password":      fmt.Sprintf("#PWD_INSTAGRAM:4:%s:%s", timestamStr, encryptedPassword),
			"contact_point": username,
			"device_id":     deviceID,
		},
		"server_params": map[string]interface{}{
			"credential_type": "password",
			"device_id":       deviceID,
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d, response: %s", resp.StatusCode, string(respBody))
	}

	raw := string(respBody)
	bearerKeyPosition := strings.Index(raw, "Bearer IGT:2:")
	if bearerKeyPosition < 0 {
		return nil, errors.New("unable to login")
	}
	key := raw[bearerKeyPosition:]
	backslashKeyPosition := strings.Index(key, "\\\\")

	token := key[13:backslashKeyPosition]

	return &model.AuthResponse{Token: token}, nil
}
