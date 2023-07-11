package model

type Post struct {
	Pk                   int64           `json:"pk"`
	ID                   string          `json:"id"`
	TextPostAppInfo      TextPostAppInfo `json:"text_post_app_info"`
	Caption              Caption         `json:"caption"`
	TakenAt              int             `json:"taken_at"`
	DeviceTimestamp      int             `json:"device_timestamp"`
	MediaType            int             `json:"media_type"`
	Code                 string          `json:"code"`
	ClientCacheKey       string          `json:"client_cache_key"`
	FilterType           int             `json:"filter_type"`
	ProductType          string          `json:"product_type"`
	OrganicTrackingToken string          `json:"organic_tracking_token"`
	ImageVersions2       struct {
		Candidates []struct {
			Width        int    `json:"width"`
			Height       int    `json:"height"`
			URL          string `json:"url"`
			ScansProfile string `json:"scans_profile"`
		} `json:"candidates"`
	} `json:"image_versions2"`
	OriginalWidth             int      `json:"original_width"`
	OriginalHeight            int      `json:"original_height"`
	VideoVersions             []string `json:"video_versions"`
	LikeCount                 int      `json:"like_count"`
	TimezoneOffset            int      `json:"timezone_offset"`
	HasLiked                  bool     `json:"has_liked"`
	LikeAndViewCountsDisabled bool     `json:"like_and_view_counts_disabled"`
	CanViewerReshare          bool     `json:"can_viewer_reshare"`
	IntegrityReviewDecision   string   `json:"integrity_review_decision"`
	TopLikers                 []string `json:"top_likers"`
	User                      User     `json:"user"`
}

type ThreadItem struct {
	Post                 Post     `json:"post"`
	LineType             string   `json:"line_type"`
	ViewRepliesCtaString string   `json:"view_replies_cta_string"`
	ShouldShowRepliesCta bool     `json:"should_show_replies_cta"`
	ReplyFacepileUsers   []string `json:"reply_facepile_users"`
	CanInlineExpandBelow bool     `json:"can_inline_expand_below"`
}

type ReplyThread struct {
	ThreadItems        []ThreadItem `json:"thread_items"`
	ThreadType         string       `json:"thread_type"`
	ShowCreateReplyCta bool         `json:"show_create_reply_cta"`
	ID                 int64        `json:"id"`
	Posts              []Post       `json:"posts"`
}

type ThreadData struct {
	ContainingThread ReplyThread   `json:"containing_thread"`
	ReplyThreads     []ReplyThread `json:"reply_threads"`
}

type ThreadDetailResponse struct {
	Data                        ThreadData        `json:"data"`
	ContainingThread            ReplyThread       `json:"containing_thread"`
	ReplyThreads                []ReplyThread     `json:"reply_threads"`
	SiblingThreads              []interface{}     `json:"sibling_threads"`
	PagingTokens                map[string]string `json:"paging_tokens"`
	DownwardsThreadWillContinue bool              `json:"downwards_thread_will_continue"`
	TargetPostReplyPlaceholder  string            `json:"target_post_reply_placeholder"`
	Status                      string            `json:"status"`
}

type User struct {
	HasAnonymousProfilePicture     bool        `json:"has_anonymous_profile_picture"`
	FanClubInfo                    FanClubInfo `json:"fan_club_info"`
	FBIDV2                         interface{} `json:"fbid_v2"`
	TransparencyProductEnabled     bool        `json:"transparency_product_enabled"`
	TextPostAppTakeABreakSetting   int         `json:"text_post_app_take_a_break_setting"`
	InteropMessagingUserFBID       int64       `json:"interop_messaging_user_fbid"`
	ShowInsightsTerms              bool        `json:"show_insights_terms"`
	AllowedCommenterType           string      `json:"allowed_commenter_type"`
	IsUnpublished                  bool        `json:"is_unpublished"`
	ReelAutoArchive                string      `json:"reel_auto_archive"`
	CanBoostPost                   bool        `json:"can_boost_post"`
	CanSeeOrganicInsights          bool        `json:"can_see_organic_insights"`
	HasOnboardedToTextPostApp      bool        `json:"has_onboarded_to_text_post_app"`
	TextPostAppJoinerNumber        int         `json:"text_post_app_joiner_number"`
	Pk                             int64       `json:"pk"`
	PKID                           string      `json:"pk_id"`
	Username                       string      `json:"username"`
	FullName                       string      `json:"full_name"`
	IsPrivate                      bool        `json:"is_private"`
	ProfilePicURL                  string      `json:"profile_pic_url"`
	AccountBadges                  []string    `json:"account_badges"`
	FeedPostReshareDisabled        bool        `json:"feed_post_reshare_disabled"`
	ShowAccountTransparencyDetails bool        `json:"show_account_transparency_details"`
	ThirdPartyDownloadsEnabled     int         `json:"third_party_downloads_enabled"`
}

type Media struct {
	TakenAt                             int64               `json:"taken_at"`
	Pk                                  int64               `json:"pk"`
	ID                                  string              `json:"id"`
	DeviceTimestamp                     int64               `json:"device_timestamp"`
	MediaType                           int                 `json:"media_type"`
	Code                                string              `json:"code"`
	ClientCacheKey                      string              `json:"client_cache_key"`
	FilterType                          int                 `json:"filter_type"`
	CanViewerReshare                    bool                `json:"can_viewer_reshare"`
	Caption                             Caption             `json:"caption"`
	ClipsTabPinnedUserIDs               []string            `json:"clips_tab_pinned_user_ids"`
	CommentInformTreatment              InformTreatment     `json:"comment_inform_treatment"`
	FundraiserTag                       FundraiserTag       `json:"fundraiser_tag"`
	SharingFrictionInfo                 SharingFrictionInfo `json:"sharing_friction_info"`
	XPostDenyReason                     string              `json:"xpost_deny_reason"`
	CaptionIsEdited                     bool                `json:"caption_is_edited"`
	OriginalMediaHasVisualReplyMedia    bool                `json:"original_media_has_visual_reply_media"`
	LikeAndViewCountsDisabled           bool                `json:"like_and_view_counts_disabled"`
	FbUserTags                          FBUserTags          `json:"fb_user_tags"`
	CanViewerSave                       bool                `json:"can_viewer_save"`
	IsInProfileGrid                     bool                `json:"is_in_profile_grid"`
	ProfileGridControlEnabled           bool                `json:"profile_grid_control_enabled"`
	FeaturedProducts                    []string            `json:"featured_products"`
	IsCommentsGifComposerEnabled        bool                `json:"is_comments_gif_composer_enabled"`
	ProductSuggestions                  []string            `json:"product_suggestions"`
	User                                User                `json:"user"`
	ImageVersions2                      ImageVersions2      `json:"image_versions2"`
	OriginalWidth                       int                 `json:"original_width"`
	OriginalHeight                      int                 `json:"original_height"`
	IsReshareOfTextPostAppMediaInIG     bool                `json:"is_reshare_of_text_post_app_media_in_ig"`
	CommentThreadingEnabled             bool                `json:"comment_threading_enabled"`
	MaxNumVisiblePreviewComments        int                 `json:"max_num_visible_preview_comments"`
	HasMoreComments                     bool                `json:"has_more_comments"`
	PreviewComments                     []Comment           `json:"preview_comments"`
	CommentCount                        int                 `json:"comment_count"`
	CanViewMorePreviewComments          bool                `json:"can_view_more_preview_comments"`
	HideViewAllCommentEntrypoint        bool                `json:"hide_view_all_comment_entrypoint"`
	Likers                              []string            `json:"likers"`
	ShopRoutingUserID                   interface{}         `json:"shop_routing_user_id"`
	CanSeeInsightsAsBrand               bool                `json:"can_see_insights_as_brand"`
	IsOrganicProductTaggingEligible     bool                `json:"is_organic_product_tagging_eligible"`
	ProductType                         string              `json:"product_type"`
	IsPaidPartnership                   bool                `json:"is_paid_partnership"`
	MusicMetadata                       interface{}         `json:"music_metadata"`
	DeletedReason                       int                 `json:"deleted_reason"`
	OrganicTrackingToken                string              `json:"organic_tracking_token"`
	TextPostAppInfo                     TextPostAppInfo     `json:"text_post_app_info"`
	IntegrityReviewDecision             string              `json:"integrity_review_decision"`
	IgMediaSharingDisabled              bool                `json:"ig_media_sharing_disabled"`
	HasSharedToFB                       int                 `json:"has_shared_to_fb"`
	IsUnifiedVideo                      bool                `json:"is_unified_video"`
	ShouldRequestAds                    bool                `json:"should_request_ads"`
	IsVisualReplyCommenterNoticeEnabled bool                `json:"is_visual_reply_commenter_notice_enabled"`
	CommercialityStatus                 string              `json:"commerciality_status"`
	ExploreHideComments                 bool                `json:"explore_hide_comments"`
	HasDelayedMetadata                  bool                `json:"has_delayed_metadata"`
}

type Caption struct {
	Pk                 string `json:"pk"`
	UserID             int64  `json:"user_id"`
	Text               string `json:"text"`
	Type               int    `json:"type"`
	CreatedAt          int64  `json:"created_at"`
	CreatedAtUTC       int64  `json:"created_at_utc"`
	ContentType        string `json:"content_type"`
	Status             string `json:"status"`
	BitFlags           int    `json:"bit_flags"`
	DidReportAsSpam    bool   `json:"did_report_as_spam"`
	ShareEnabled       bool   `json:"share_enabled"`
	User               User   `json:"user"`
	IsCovered          bool   `json:"is_covered"`
	IsRankedComment    bool   `json:"is_ranked_comment"`
	MediaID            int64  `json:"media_id"`
	PrivateReplyStatus int    `json:"private_reply_status"`
}

type InformTreatment struct {
	ShouldHaveInformTreatment bool   `json:"should_have_inform_treatment"`
	Text                      string `json:"text"`
	URL                       string `json:"url"`
	ActionType                string `json:"action_type"`
}

type FundraiserTag struct {
	HasStandaloneFundraiser bool `json:"has_standalone_fundraiser"`
}

type SharingFrictionInfo struct {
	ShouldHaveSharingFriction bool   `json:"should_have_sharing_friction"`
	BloksAppURL               string `json:"bloks_app_url"`
	SharingFrictionPayload    string `json:"sharing_friction_payload"`
}

type FBUserTags struct {
	In []string `json:"in"`
}

type ImageVersions2 struct {
	Candidates []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"candidates"`
}

type Comment struct {
	Pk              int64  `json:"pk"`
	UserID          int64  `json:"user_id"`
	Text            string `json:"text"`
	Type            int    `json:"type"`
	CreatedAt       int64  `json:"created_at"`
	CreatedAtUTC    int64  `json:"created_at_utc"`
	ContentType     string `json:"content_type"`
	Status          string `json:"status"`
	BitFlags        int    `json:"bit_flags"`
	DidReportAsSpam bool   `json:"did_report_as_spam"`
	ShareEnabled    bool   `json:"share_enabled"`
	User            User   `json:"user"`
}

type TextPostAppInfo struct {
	IsPostUnavailable     bool        `json:"is_post_unavailable"`
	IsReply               bool        `json:"is_reply"`
	ReplyToAuthor         interface{} `json:"reply_to_author"`
	DirectReplyCount      int         `json:"direct_reply_count"`
	SelfThreadCount       int         `json:"self_thread_count"`
	ReplyFacepileUsers    []User      `json:"reply_facepile_users"`
	LinkPreviewAttachment interface{} `json:"link_preview_attachment"`
	CanReply              bool        `json:"can_reply"`
	ReplyControl          string      `json:"reply_control"`
	HushInfo              interface{} `json:"hush_info"`
	ShareInfo             ShareInfo   `json:"share_info"`
}

type ShareInfo struct {
	CanRepost          bool `json:"can_repost"`
	IsRepostedByViewer bool `json:"is_reposted_by_viewer"`
	CanQuotePost       bool `json:"can_quote_post"`
}

type FanClubInfo struct {
	FanClubID                            int64       `json:"fan_club_id"`
	FanClubName                          string      `json:"fan_club_name"`
	IsFanClubReferralEligible            interface{} `json:"is_fan_club_referral_eligible"`
	FanConsiderationPageRevampEligiblity interface{} `json:"fan_consideration_page_revamp_eligiblity"`
	IsFanClubGiftingEligible             interface{} `json:"is_fan_club_gifting_eligible"`
	SubscriberCount                      interface{} `json:"subscriber_count"`
	ConnectedMemberCount                 interface{} `json:"connected_member_count"`
	AutosaveToExclusiveHighlight         interface{} `json:"autosave_to_exclusive_highlight"`
	HasEnoughSubscribersForSSC           interface{} `json:"has_enough_subscribers_for_ssc"`
}
