package threads

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"

	"github.com/dwarvesf/go-threads/model"
	"github.com/dwarvesf/go-threads/private"
)

type PrivateAPI interface {
	Auth() (*model.AuthResponse, error)
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

func NewPrivateAPIClient(uID int, username string, password string, androidDeviceID, token string) (PrivateAPI, error) {
	if uID == 0 {
		userID, err := GetUserByUsername(username)
		if err != nil {
			return nil, err
		}

		uID = userID
	}

	ins := &private.PrivateAPI{
		UserID:          uID,
		Username:        username,
		Password:        password,
		TimezoneOffset:  -14400,
		AndroidDeviceID: androidDeviceID,
		APIToken:        token,
	}

	if token == "" {
		ins.Auth()
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

// GetUserByUsername get user by username
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

	respBody, err := httputil.DumpResponse(resp, true)
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
