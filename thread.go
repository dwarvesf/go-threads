package threads

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/dwarvesf/go-threads/model"
	"github.com/dwarvesf/go-threads/private"
)

// PrivateAPI interface for private API
type PrivateAPI interface {
	CreatePost(content model.CreatePostRequest) (*model.CreatePostResponse, error)
	GetUserFollowers(id int) (*model.UserFollowersResponse, error)
	GetFollowers() (*model.UserFollowersResponse, error)
	GetFollowing() (*model.UserFollowingResponse, error)
	GetThreadByID(ID string) (*model.ThreadDetailResponse, error)
	SearchUser(query string) (*model.UserResponse, error)
	LikeThread(id string) (map[string]interface{}, error)
	UnLikeThread(id string) (map[string]interface{}, error)
	FollowUser(id int) (*model.FollowUserResponse, error)
	UnFollowUser(id int) (*model.FollowUserResponse, error)
}

// NewPrivateAPIClient new api client for private API
func NewPrivateAPIClient(cfg *Config) (PrivateAPI, error) {
	if err := cfg.ReadyCheck(); err != nil {
		return nil, err
	}

	ins := &private.PrivateAPI{
		UserID:             cfg.UserID,
		Username:           cfg.Username,
		Password:           cfg.Password,
		TimezoneOffset:     cfg.TimezoneOffset,
		DeviceID:           cfg.DeviceID,
		DeviceManufacturer: cfg.DeviceManufacturer,
		DeviceModel:        cfg.DeviceModel,
		DeviceOsVersion:    cfg.DeviceOsVersion,
		DeviceOsRelease:    cfg.DeviceOsRelease,
		APIToken:           cfg.APIToken,
	}

	return ins, nil
}

// InitAPIClient new api client for private API
func InitAPIClient(cfgFn ...ConfigFn) (PrivateAPI, error) {
	cfg, err := InitConfig(cfgFn...)
	if err != nil {
		return nil, err
	}
	if err := cfg.ReadyCheck(); err != nil {
		return nil, err
	}

	ins := &private.PrivateAPI{
		UserID:             cfg.UserID,
		Username:           cfg.Username,
		Password:           cfg.Password,
		TimezoneOffset:     cfg.TimezoneOffset,
		DeviceID:           cfg.DeviceID,
		DeviceManufacturer: cfg.DeviceManufacturer,
		DeviceModel:        cfg.DeviceModel,
		DeviceOsVersion:    cfg.DeviceOsVersion,
		DeviceOsRelease:    cfg.DeviceOsRelease,
		APIToken:           cfg.APIToken,
	}

	return ins, nil
}

var fetchHTMLHeaders = map[string]string{
	"Authority":                 "www.threads.net",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"Accept-Language":           "en-US,en;q=0.9",
	"Cache-Control":             "no-cache",
	"Content-Type":              "application/x-www-form-urlencoded",
	"Origin":                    "https://www.threads.net",
	"Pragma":                    "no-cache",
	"Referer":                   "https://www.instagram.com",
	"Sec-Fetch-Dest":            "document",
	"Sec-Fetch-Mode":            "navigate",
	"Sec-Fetch-Site":            "cross-site",
	"Sec-Fetch-User":            "?1",
	"Upgrade-Insecure-Requests": "1",
}

// GetUserByUsername get user by username via instagram api
func GetUserByUsername(username string) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.instagram.com/%s", username), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	for key, value := range fetchHTMLHeaders {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	userIDKeyPattern := regexp.MustCompile(`"user_id":"(\d+)",`)
	userIDKeyMatch := userIDKeyPattern.FindStringSubmatch(string(respBody))
	if userIDKeyMatch == nil {
		return 0, fmt.Errorf("failed to find user ID in the response")
	}

	userID := userIDKeyMatch[1]

	return strconv.Atoi(userID)
}
