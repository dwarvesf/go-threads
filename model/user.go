package model

type UserProfile struct {
	HasAnonymousProfilePicture bool         `json:"has_anonymous_profile_picture"`
	FollowerCount              int          `json:"follower_count"`
	MediaCount                 int          `json:"media_count"`
	FollowingCount             int          `json:"following_count"`
	FollowingTagCount          int          `json:"following_tag_count"`
	FbidV2                     string       `json:"fbid_v2"`
	HasOnboardedToTextPostApp  bool         `json:"has_onboarded_to_text_post_app"`
	ShowTextPostAppBadge       bool         `json:"show_text_post_app_badge"`
	TextPostAppJoinerNumber    int          `json:"text_post_app_joiner_number"`
	ShowIGAppSwitcherBadge     bool         `json:"show_ig_app_switcher_badge"`
	Pk                         int          `json:"pk"`
	PkID                       string       `json:"pk_id"`
	Username                   string       `json:"username"`
	FullName                   string       `json:"full_name"`
	IsPrivate                  bool         `json:"is_private"`
	IsVerified                 bool         `json:"is_verified"`
	ProfilePicID               string       `json:"profile_pic_id"`
	ProfilePicURL              string       `json:"profile_pic_url"`
	HasOptEligibleShop         bool         `json:"has_opt_eligible_shop"`
	AccountBadges              []string     `json:"account_badges"`
	ThirdPartyDownloadsEnabled int          `json:"third_party_downloads_enabled"`
	UnseenCount                int          `json:"unseen_count"`
	FriendshipStatus           FriendStatus `json:"friendship_status"`
	LatestReelMedia            int          `json:"latest_reel_media"`
	ShouldShowCategory         bool         `json:"should_show_category"`
}

type FriendStatus struct {
	Following               bool `json:"following"`
	IsPrivate               bool `json:"is_private"`
	IncomingRequest         bool `json:"incoming_request"`
	OutgoingRequest         bool `json:"outgoing_request"`
	TextPostAppPreFollowing bool `json:"text_post_app_pre_following"`
	IsBestie                bool `json:"is_bestie"`
	IsRestricted            bool `json:"is_restricted"`
	IsFeedFavorite          bool `json:"is_feed_favorite"`
}

type UserResponse struct {
	NumResults int           `json:"num_results"`
	Users      []UserProfile `json:"users"`
	HasMore    bool          `json:"has_more"`
	RankToken  string        `json:"rank_token"`
	Status     string        `json:"status"`
}
