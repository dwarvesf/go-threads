package public

import "net/http"

const ThreadsAPIURL = "https://www.threads.net/api/graphql"

type PublicAPI struct {
	APPToken string
}

func prepareHeaders(req *http.Request, token string) *http.Request {
	for k, v := range defaultHeaders(token) {
		req.Header.Set(k, v)
	}

	return req
}

func defaultHeaders(token string) map[string]string {
	return map[string]string{
		"Authority":       "www.threads.net",
		"Accept":          "*/*",
		"Accept-Language": "en-US,en;q=0.9",
		"Cache-Control":   "no-cache",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Origin":          "https://www.threads.net",
		"Pragma":          "no-cache",
		"Sec-Fetch-Site":  "same-origin",
		"X-ASBD-ID":       "129477",
		"X-FB-LSD":        token,
		"X-IG-App-ID":     "238260118697367",
	}
}
