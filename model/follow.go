package model

type UserSummary struct {
	HasAnonymousProfilePicture bool        `json:"has_anonymous_profile_picture"`
	FBIDV2                     interface{} `json:"fbid_v2"`
	HasOnboardedToTextPostApp  bool        `json:"has_onboarded_to_text_post_app"`
	TextPostAppJoinerNumber    int         `json:"text_post_app_joiner_number"`
	PK                         int64       `json:"pk"`
	PKID                       string      `json:"pk_id"`
	Username                   string      `json:"username"`
	FullName                   string      `json:"full_name"`
	IsPrivate                  bool        `json:"is_private"`
	IsVerified                 bool        `json:"is_verified"`
	ProfilePicID               string      `json:"profile_pic_id"`
	ProfilePicURL              string      `json:"profile_pic_url"`
	AccountBadges              []string    `json:"account_badges"`
	IsPossibleScammer          bool        `json:"is_possible_scammer"`
	ThirdPartyDownloadsEnabled int         `json:"third_party_downloads_enabled"`
	IsPossibleBadActor         struct {
		IsPossibleScammer      bool `json:"is_possible_scammer"`
		IsPossibleImpersonator struct {
			IsUnconnectedImpersonator bool `json:"is_unconnected_impersonator"`
		} `json:"is_possible_impersonator"`
	} `json:"is_possible_bad_actor"`
	LatestReelMedia int `json:"latest_reel_media"`
}

type UserFollowersResponse struct {
	Users                      []UserSummary  `json:"users"`
	BigList                    bool           `json:"big_list"`
	PageSize                   int            `json:"page_size"`
	Groups                     []interface{}  `json:"groups"`
	MoreGroupsAvailable        bool           `json:"more_groups_available"`
	FriendRequests             map[string]int `json:"friend_requests"`
	HasMore                    bool           `json:"has_more"`
	ShouldLimitListOfFollowers bool           `json:"should_limit_list_of_followers"`
	Status                     string         `json:"status"`
}

type UserFollowingResponse struct {
	Users                      []UserSummary `json:"users"`
	BigList                    bool          `json:"big_list"`
	PageSize                   int           `json:"page_size"`
	HasMore                    bool          `json:"has_more"`
	ShouldLimitListOfFollowers bool          `json:"should_limit_list_of_followers"`
	Status                     string        `json:"status"`
}

type FriendshipStatus struct {
	Following               bool `json:"following"`
	FollowedBy              bool `json:"followed_by"`
	Blocking                bool `json:"blocking"`
	Muting                  bool `json:"muting"`
	IsPrivate               bool `json:"is_private"`
	IncomingRequest         bool `json:"incoming_request"`
	OutgoingRequest         bool `json:"outgoing_request"`
	TextPostAppPreFollowing bool `json:"text_post_app_pre_following"`
	IsBestie                bool `json:"is_bestie"`
	IsRestricted            bool `json:"is_restricted"`
	IsFeedFavorite          bool `json:"is_feed_favorite"`
	IsEligibleToSubscribe   bool `json:"is_eligible_to_subscribe"`
}

type FollowUserResponse struct {
	FriendshipStatus  FriendshipStatus `json:"friendship_status"`
	PreviousFollowing bool             `json:"previous_following"`
	Status            string           `json:"status"`
}
