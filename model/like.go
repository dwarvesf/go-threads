package model

type LikersResponse struct {
	Data struct {
		Likers struct {
			Users []UserInfo `json:"users"`
		} `json:"likers"`
	} `json:"data"`
	Extensions struct {
		IsFinal bool `json:"is_final"`
	} `json:"extensions"`
}

type UserInfo struct {
	PK                     string      `json:"pk"`
	FullName               string      `json:"full_name"`
	ProfilePicURL          string      `json:"profile_pic_url"`
	FollowerCount          int         `json:"follower_count"`
	IsVerified             bool        `json:"is_verified"`
	Username               string      `json:"username"`
	ProfileContextFacepile interface{} `json:"profile_context_facepile_users"`
	ID                     interface{} `json:"id"`
}
