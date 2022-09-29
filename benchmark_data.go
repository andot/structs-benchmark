package structsbenchmark_test

import "strings"

type StringArray []string

func (v StringArray) String() string {
	return strings.Join(v, "\n")
}

type Response struct {
	ID                               string                          `structs:"response_id"`
	AdID                             string                          `structs:"response_ad_id"`
	CreativeID                       string                          `structs:"response_creative_id"`
	Int64CreativeID                  int64                           `structs:"response_int64_creative_id"`
	AdvertiserCreativeID             string                          `structs:"response_advertiser_creative_id"`
	HTMLSnippet                      string                          `structs:"response_html_snippet,omitempty"`
	HTMLURL                          string                          `structs:"response_html_url,omitempty"`
	ImageURL                         string                          `structs:"response_image_url,omitempty"`
	ImageWidth                       int64                           `structs:"response_image_width"`
	ImageHeight                      int64                           `structs:"response_image_height"`
	CoverURL                         string                          `structs:"response_cover_url,omitempty"`
	IconURL                          string                          `structs:"response_icon_url,omitempty"`
	IconWidth                        int64                           `structs:"response_icon_width,omitempty"`
	IconHeight                       int64                           `structs:"response_icon_height,omitempty"`
	LogoURL                          string                          `structs:"response_logo_url,omitempty"`
	Width                            int64                           `structs:"response_width"`
	Height                           int64                           `structs:"response_height"`
	AdvertisingAppName               string                          `structs:"response_advertising_app_name,omitempty"`
	AdvertisingAppBundle             string                          `structs:"response_advertising_app_bundle,omitempty"`
	AdvertisingAppVersion            string                          `structs:"response_advertising_app_version,omitempty"`
	TargetURL                        string                          `structs:"response_target_url"`
	Deeplink                         string                          `structs:"response_deeplink"`
	Scheme                           string                          `structs:"response_scheme,omitempty"`
	ClickTrackers                    StringArray                     `structs:"response_click_trackers,string"`
	ImpressionTrackers               StringArray                     `structs:"response_impression_trackers,string"`
	DownloadStartTrackers            StringArray                     `structs:"response_download_start_trackers,string"`
	DownloadEndedTrackers            StringArray                     `structs:"response_download_ended_trackers,string"`
	AppInstalledTrackers             StringArray                     `structs:"response_app_installed_trackers,string"`
	RefreshInterval                  int64                           `structs:"response_refresh_interval"`
	Title                            string                          `structs:"response_title"`
	SubTitle                         string                          `structs:"response_sub_title"`
	ButtonText                       string                          `structs:"response_button_text"`
	Description                      string                          `structs:"response_description"`
	FileSize                         int64                           `structs:"response_file_size"`
	PriceInCent                      float64                         `structs:"response_price_in_cent"`
	IsBiddingAdvertiser              bool                            `structs:"response_is_bidding_advertiser"`
	BottomPrice                      float64                         `structs:"response_bottom_price"`
	PriceReplacer                    func(float64) *strings.Replacer `structs:"-"`
	AdvertiserID                     int64                           `structs:"response_advertising_advertiser_id"`
	AdvertiserIDFromMonetizer        string                          `structs:"response_advertising_advertiser_id_from_monetizer"`
	AdvertiserName                   string                          `structs:"response_advertising_advertiser_name"`
	ExtendedParameters               StringArray                     `structs:"response_extended_parameters,string"`
	Video                            *ResponseVideo                  `structs:",flatten"`
	MonetizerAppAdUnitToken          string                          `structs:"response_monetizer_app_ad_unit_token"`
	DeeplinkAppNotInstalledTrackers  StringArray                     `structs:"response_deeplink_app_not_installed_trackers,string"`
	DeeplinkAppInstalledTrackers     StringArray                     `structs:"response_deeplink_app_installed_trackers,string"`
	DeeplinkAppInvokeFailedTrackers  StringArray                     `structs:"response_deeplink_app_invoke_failed_trackers,string"`
	DeeplinkAppInvokeSuccessTrackers StringArray                     `structs:"response_deeplink_app_invoke_success_trackers,string"`
	ImagesForNativeResponse          StringArray                     `structs:"response_images,string"`
	WinNoticeTrackers                StringArray                     `structs:"response_win_notice_trackers,string"`
	WinNoticeTracker                 string                          `structs:"response_win_notice_tracker"`
	LossNoticeTracker                string                          `structs:"response_loss_notice_tracker"`
	AdvertiserDomain                 string                          `structs:"response_advertiser_domain"`
	PrivacyURL                       string                          `structs:"response_privacy_url"`
	PermissionURL                    string                          `structs:"response_permission_url"`
	PermissionTitles                 StringArray                     `structs:"response_permission_titles,string"`
	PermissionDescriptions           StringArray                     `structs:"response_permission_descriptions,string"`
	NoFillingExpiration              uint32                          `structs:"-"`
}

type ResponseVideo struct {
	URL               string      `structs:"response_video_url"`
	PlayDuration      int64       `structs:"response_video_play_duration"`
	Size              int64       `structs:"response_video_size"`
	Width             int64       `structs:"response_video_width"`
	Height            int64       `structs:"response_video_height"`
	BitRate           int64       `structs:"response_video_bit_rate"`
	PlayStartTrackers StringArray `structs:"response_video_play_start_trackers,string"`
	PlayBreakTrackers StringArray `structs:"response_video_play_break_trackers,string"`
	PlayEndedTrackers StringArray `structs:"response_video_play_ended_trackers,string"`
}

func NewResponse() *Response {
	return &Response{
		ID:                               "1",
		AdID:                             "1234",
		CreativeID:                       "fjadsltueioaaja",
		Int64CreativeID:                  12458594305832,
		AdvertiserCreativeID:             "3389324",
		HTMLSnippet:                      "",
		ImageURL:                         "https://camo.githubusercontent.com/4ba61d65b98f20032813be8e1bb9764bb5771fdc3947ae9e1537b228ebc50259/687474703a2f2f696d672e736869656c64732e696f2f636f766572616c6c732f616e646f742f737472756374732e7376673f7374796c653d666c61742d737175617265",
		ImageWidth:                       1080,
		ImageHeight:                      1920,
		CoverURL:                         "http://img.shields.io/coveralls/andot/structs.svg?style=flat-square",
		IconURL:                          "http://img.shields.io/travis/andot/structs.svg?style=flat-square",
		IconWidth:                        100,
		IconHeight:                       100,
		Width:                            1080,
		Height:                           1920,
		AdvertisingAppName:               "structs",
		AdvertisingAppBundle:             "structs.benchmark",
		AdvertisingAppVersion:            "1.3.0",
		TargetURL:                        "https://travis-ci.org/andot/structs",
		Deeplink:                         "https://coveralls.io/r/andot/structs",
		ClickTrackers:                    []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		ImpressionTrackers:               []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		DownloadStartTrackers:            []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		DownloadEndedTrackers:            []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		AppInstalledTrackers:             []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		DeeplinkAppNotInstalledTrackers:  []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		DeeplinkAppInstalledTrackers:     []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		DeeplinkAppInvokeFailedTrackers:  []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		DeeplinkAppInvokeSuccessTrackers: []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		ImagesForNativeResponse:          []string{"http://img.shields.io/coveralls/andot/structs.svg?style=flat-square", "http://img.shields.io/travis/andot/structs.svg?style=flat-square"},
		Video: &ResponseVideo{
			URL:               "http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square",
			PlayDuration:      30,
			Size:              12984201823,
			Width:             1080,
			Height:            1920,
			BitRate:           12370312,
			PlayStartTrackers: []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
			PlayBreakTrackers: []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
			PlayEndedTrackers: []string{"https://travis-ci.org/andot/structs", "https://coveralls.io/r/andot/structs"},
		},
	}
}
