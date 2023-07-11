package model

type CreatePostRequest struct {
	Caption string
}

type CreatePostResponse struct {
	Media Media `json:"media,omitempty"`
}
