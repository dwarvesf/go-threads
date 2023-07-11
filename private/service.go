package private

import (
	"net/http"
)

type PrivateAPI struct {
	UserID          int
	Username        string
	Password        string
	TimezoneOffset  int
	AndroidDeviceID string
	APIToken        string
}

func generateAuthHeader(token string) map[string]string {
	return map[string]string{
		"Authorization":  "Bearer IGT:2:" + token,
		"User-Agent":     "Barcelona 289.0.0.77.109 Android",
		"Sec-Fetch-Site": "same-origin",
		"Content-Type":   "application/x-www-form-urlencoded; charset=UTF-8",
	}
}

func updateRequestHeader(token string, req *http.Request) *http.Request {
	headers := generateAuthHeader(token)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req
}
