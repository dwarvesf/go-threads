package threads

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"sync"

	"github.com/dwarvesf/go-threads/model"
	"github.com/dwarvesf/go-threads/public"
)

var instance PublicAPI
var once sync.Once

func getPublicAPIClient() PublicAPI {
	once.Do(func() {
		c, err := newPublicAPIClient()
		if err != nil {
			fmt.Println("unable to init public API: " + err.Error())
		}
		instance = c
	})
	return instance
}

func newPublicAPIClient() (*public.PublicAPI, error) {
	token, err := getThreadsAPIToken()
	if err != nil {
		return nil, err
	}

	return &public.PublicAPI{
		APPToken: token,
	}, nil
}

// PublicAPI interface for Public API
type PublicAPI interface {
	GetThreadLikers(ID int) (*model.LikersResponse, error)
}

func getThreadsAPIToken() (string, error) {
	url := "https://www.instagram.com/instagram"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	for k, v := range fetchHTMLHeaders {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`LSD",\[\],{"token":"(.*?)"},\d+]`)
	match := re.FindStringSubmatch(string(body))
	if len(match) < 2 {
		return "", fmt.Errorf("failed to extract token")
	}

	token := match[1]

	return token, nil
}

func GetThreadLikers(ID int) (*model.LikersResponse, error) {
	return getPublicAPIClient().GetThreadLikers(ID)
}
