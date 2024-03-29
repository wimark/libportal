package libportal

import (
	"sort"
	"strings"
	"time"
)

// PortalAuthObject is struct for store every auth attempt
// (Deprecated)
type PortalAuthObject struct {
	Timestamp int64 `json:"timestamp" bson:"timestamp"`

	CPE       string `json:"cpe_id" bson:"cpe_id"`
	Ip        string `json:"client_ip" bson:"client_ip"`
	Useragent string `json:"useragent" bson:"useragent"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// PortalClientSession is struct for store portal client
type PortalClientSession struct {
	ID string `json:"id" bson:"_id"`

	// identity
	MAC     string `json:"mac" bson:"mac"`
	WLAN    string `json:"wlan_id" bson:"wlan_id"`
	CPE     string `json:"cpe" bson:"cpe"`
	NAS     string `json:"nas" bson:"nas"`
	LocID   string `json:"loc_id" bson:"loc_id"`
	IP      string `json:"ip" bson:"ip"`
	Hash    string `json:"hash" bson:"hash"`
	Segment string `json:"segment" bson:"segment"`

	// link with profile and account
	Profile string `json:"profile" bson:"profile"`
	Account string `json:"account" bson:"account"`

	// creation time (index needs)
	Create time.Time `json:"create" bson:"create"`

	// from create to end of internet + block
	CreateAt int64 `json:"create_at" bson:"create_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	// real internet part
	StartAt  int64 `json:"start_at" bson:"start_at"`
	StopAt   int64 `json:"stop_at" bson:"stop_at"`
	Duration int64 `json:"duration" bson:"duration"`

	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`
	Auth          []PortalAuthObject  `json:"auth" bson:"auth"`

	State       PortalUserState `json:"state" bson:"state"`
	AuthenState PortalAuthenticationState
	AuthState   PortalAuthorizationState

	Authentication PortalAuthenticationData `json:"authentication" bson:"authentication"`
	Authorization  PortalAuthorizationData  `json:"authorization" bson:"authorization"`

	// data from UserAgent
	UA     UserAgent `json:"ua" bson:"ua"`
	Locale string    `json:"locale" bson:"locale"`

	// will be deprecated
	Status   string `json:"status" bson:"status"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`

	// "removal" of identification and authorization of clients
	IsDeleted bool `json:"is_deleted" bson:"is_deleted"`

	// Cost sms from sms.ru
	Cost float64 `json:"cost" bson:"cost"`

	SocialNetwork map[string]AccountFromSocialNetwork `json:"social_network" bson:"social_network"`

	AcctSessionID string `json:"acct_session_id" bson:"acct_session_id"`
}

// PortalClientAuthentication struct for store portal client authen
// client could be authenticated in one profile
type PortalClientAuthentication struct {
	Id string `json:"id" bson:"_id"`

	// client identification
	MAC     string `json:"mac" bson:"mac"`
	Profile string `json:"profile" bson:"profile"`

	WLAN string `json:"wlan_id" bson:"wlan_id"` // not actually needed

	// data with identity and authentication info
	Data PortalAuthenticationData `json:"data" bson:"data"`

	Create time.Time `json:"create" bson:"create"`

	// valid time
	CreateAt int64 `json:"create_at" bson:"create_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	// "removal" of identification and authorization of clients
	IsDeleted bool `json:"is_deleted" bson:"is_deleted"`
}

// PortalCondition struct for apply profile condition
type PortalCondition struct {
	WLAN  []string `json:"wlan" bson:"wlan"`
	CPE   []string `json:"cpe" bson:"cpe"`
	NasID []string `json:"nas_id" bson:"nas_id"`
	LocID []string `json:"loc_id" bson:"loc_id"`
}

// Empty func to check struct for empty
func (pc *PortalCondition) Empty() bool {
	return len(pc.WLAN) == 0 && len(pc.CPE) == 0 && len(pc.NasID) == 0 && len(pc.LocID) == 0
}

// PortalConditionRequest request for condition of profile
type PortalConditionRequest struct {
	WLAN  string `json:"wlan,omitempty" bson:"wlan"`
	CPE   string `json:"cpe,omitempty" bson:"cpe"`
	NasID string `json:"nas_id,omitempty" bson:"nas_id"`
	LocID string `json:"loc_id,omitempty" bson:"loc_id"`
}

// PortalSessionConfig struct for flexible session config
type PortalSessionConfig struct {

	// session timeout -- 30 min as example
	Timeout int64 `json:"timeout" bson:"timeout"`

	// timeout to remember
	AuthTimeout int64 `json:"auth_timeout" bson:"auth_timeout"`

	// block after using timeout for (example: 3600 is for seconds online)
	BlockAfterTimeout int64 `json:"block_after" bson:"block_after"`

	// expiration of block after (example: 7200 is for seconds to nullify block_after)
	BlockExpireTimeout int64 `json:"block_expire" bson:"block_expire"`

	// traffic limit
	DownloadLimit int `json:"download_limit" bson:"download_limit"`

	// max number
	MaxSessions int `json:"max_sessions" bson:"max_sessions"`

	// DEPRECATED field for block timeout
	BlockTimeout int64 `json:"block_timeout" bson:"block_timeout"`
}

// PortalMSISDNConfig config of possible CC and NDC prefixes in MSISDN
type PortalMSISDNConfig struct {
	// example: {"992": ["90","92"...]}
	Prefix map[string][]string `json:"prefix" bson:"prefix"`
	PrMap  []string            `json:"-" bson:"map"`
}

// Map (PortalMSISDNConfig) func to Map DEF / CC codes
func (p *PortalMSISDNConfig) Map() {
	p.PrMap = []string{}
	for k, v := range p.Prefix {
		if len(v) == 0 {
			p.PrMap = append(p.PrMap, k)
		} else {
			for _, vv := range v {
				p.PrMap = append(p.PrMap, k+vv)
			}
		}
	}
}

func (p PortalMSISDNConfig) Check(phone string) bool {
	if len(p.PrMap) > 0 && len(phone) > 0 {
		for _, v := range p.PrMap {
			if strings.HasPrefix(phone, v) {
				return true
			}
		}
		return false
	}
	return true
}

type PortalAuthenticationData struct {
	Type PortalAuthenticationType `json:"type" bson:"type"`

	Identity string `json:"identity" bson:"identity"`
	Password string `json:"password" bson:"password"`

	Remember      int64 `json:"remember" bson:"remember"`
	PushAgreement bool  `json:"push_agree" bson:"push_agree"`

	// additional data if needed (for example full ESIA string)
	Data string `json:"data" bson:"data"`
}

// PortalAuthenticationConfig config of possible authentications
type PortalAuthenticationConfig struct {
	Id string `json:"id" bson:"_id"`

	Enable bool                     `json:"enable" bson:"enable"`
	Type   PortalAuthenticationType `json:"type" bson:"type"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	Remember  int64 `json:"remember" bson:"remember"`
	OPTLength int   `json:"otp_length" bson:"otp_length"`

	PassGateway string `json:"-" bson:"pass_gateway"`

	PortalHTTPRequest PortalHTTPRequest `json:"portal_http_request" bson:"portal_http_request"`

	Validate struct {
		Regex    string `json:"regex" bson:"regex"`
		ErrorMsg string `json:"error" bson:"error"`
	} `json:"validate" bson:"validate"`
}

type PortalHTTPRequest struct {
	Method      string `json:"method" bson:"method"`
	Url         string `json:"url" bson:"url"`
	ContentType string `json:"content_type" bson:"content_type"`
	Payload     string `json:"payload" bson:"payload"`
}

type PortalAuthorizationData struct {
	ConfigId string `json:"config_id" bson:"config_id"`

	Type  PortalAuthorizationType  `json:"type" bson:"type"`
	State PortalAuthorizationState `json:"state" bson:"state"`

	Config PortalSessionConfig `json:"config" bson:"config"`

	// will be deprecated
	Ads []PortalAd `json:"ads" bson:"-"`

	AdsIDs         []string `json:"ads_ids" bson:"ads_ids"`
	AdsToWatch     int      `json:"ads_to_watch" bson:"ads_to_watch"`
	EnableRotation bool     `json:"enable_rotation" bson:"enable_rotation"`
	EnableAdFollow bool     `json:"enable_ad_follow" bson:"enable_ad_follow"`

	RedirectURL string `json:"redirect_url" bson:"redirect_url"`

	Notification PortalNotification `json:"notification" bson:"notification"`

	//SocialNetworkConfig SN auth setting
	SocialNetworkConfig PortalAuthorizationWithSN `json:"social_network_config" bson:"social_network_config"`

	SkipFinalPage bool `json:"skip_final_page" bson:"skip_final_page"`
	Transferable  bool `json:"transferable" bson:"transferable"`
}

type PortalAuthorizationConfig struct {
	Id string `json:"id" bson:"_id"`

	Enable bool                    `json:"enable" bson:"enable"`
	Type   PortalAuthorizationType `json:"type" bson:"type"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// header and info for showing on captive portal page
	Header string `json:"header" bson:"header"`
	Info   string `json:"info" bson:"info"`

	// advertisements for user
	AdsIDs         []string `json:"ads_ids" bson:"ads_ids"`
	AdsToWatch     int      `json:"ads_to_watch" bson:"ads_to_watch"`
	EnableRotation bool     `json:"enable_rotation" bson:"enable_rotation"`
	EnableAdFollow bool     `json:"enable_ad_follow" bson:"enable_ad_follow"`

	// session configuration
	Config PortalSessionConfig `json:"config" bson:"config"`

	// landing page redirect URL
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`

	// notification config
	Notification PortalNotification `json:"notification" bson:"notification"`

	//SocialNetworkConfig SN auth setting
	SocialNetworkConfig PortalAuthorizationWithSN `json:"social_network_config" bson:"social_network_config"`

	// skip final page (with data about how much to internet)
	SkipFinalPage bool `json:"skip_final_page" bson:"skip_final_page"`

	// is authorisation transferable on NAS
	// thus it could be divided to small chunks on NAS
	Transferable bool `json:"transferable" bson:"transferable"`

	// DEPRECATED
	Ads []PortalAd `json:"ads" bson:"ads"`
}

func (p *PortalAuthorizationConfig) SortAd() {
	sort.Slice(p.Ads, func(i, j int) bool {
		return p.Ads[i].Priority > p.Ads[j].Priority
	})
}

// PortalNotification struct for send notification with text before session stop
type PortalNotification struct {
	Enable            bool   `json:"enable" bson:"enable"`
	Text              string `json:"text" bson:"text"`
	SecondsBeforeStop int    `json:"seconds_before_stop" bson:"seconds_before_stop"`
}

// PortalProfile to link portal and it's config
type PortalProfile struct {
	Id string `json:"id" bson:"_id"`

	ShortID     string `json:"short_id" bson:"short_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// condition to check
	// TODO: check profile special API to find what profile will be for condition
	Condition PortalCondition `json:"condition" bson:"condition"`

	// authentication types
	Authentications    []PortalAuthenticationConfig `json:"authentications" bson:"authentications"`
	AuthenticationSkip bool                         `json:"authn_skip" bson:"authn_skip"`

	// available MSISDN prefixes --
	MSISDNConfig PortalMSISDNConfig `json:"msisdn_config" bson:"msisdn_config"`

	// authorization types
	Authorizations       []PortalAuthorizationConfig `json:"authorizations" bson:"authorizations"`
	AuthorizationDefault string                      `json:"authz_default" bson:"authz_default"`

	// limits for authens and auths per page
	AuthenticationLimit int `json:"authentication_limit" bson:"authentication_limit"`
	AuthorizationLimit  int `json:"authorization_limit" bson:"authorization_limit"`

	//AccessList list to pass special user by code
	AccessList map[string]bool `json:"access_list" bson:"access_list"`
	//WhiteList to pass user with out auth
	WhiteList map[string]bool `json:"white_list" bson:"white_list"`
	//BlackList is list to block user
	BlackList map[string]bool `json:"black_list" bson:"black_list"`

	// default session configuration (timeout and block timeout)
	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`

	// UTC diff (plus or minus from UTC time)
	UTCDiff int `json:"utc_diff" bson:"utc_diff"`

	MACsLimit int `json:"macs_limit" bson:"macs_limit"`

	// to payments and payments system integration
	AllowBalance   bool     `json:"allow_balance" bson:"allow_balance"`
	PaymentSystems []string `json:"payment_systems" bson:"payment_systems"`

	// settings for terms of service (no locale)
	TermsOfService struct {
		Enable   bool   `json:"enable" bson:"enable"`
		External bool   `json:"external" bson:"external"`
		FileURL  string `json:"file_url" bson:"file_url"`
		Text     string `json:"text" bson:"text"`
	} `json:"terms_of_service" bson:"terms_of_service"`

	// messages for OTP codes / vouchers and subscriptions
	PushText struct {
		OTP          string `json:"otp" bson:"otp"`
		Voucher      string `json:"voucher" bson:"voucher"`
		Subscription string `json:"subscription" bson:"subscription"`
	} `json:"push_text" bson:"push_text"`

	// Block page for Random (Local Assigned) MAC Addresses
	RandomMACBlockPage bool `json:"random_mac_block_page" bson:"random_mac_block_page"`

	// settings for PushAgreement (no locale)
	PushAgreement struct {
		Enable   bool   `json:"enable" bson:"enable"`
		External bool   `json:"external" bson:"external"`
		FileURL  string `json:"file_url" bson:"file_url"`
		Text     string `json:"text" bson:"text"`
	} `json:"push_agreement" bson:"push_agreement"`

	// service private area
	EnableLineCabinet bool `json:"enable_line_cabinet" bson:"enable_line_cabinet"`
}

func (p *PortalProfile) SortAd() {
	for i := range p.Authorizations {
		p.Authorizations[i].SortAd()
	}
}

type pageColor struct {
	Main       string `json:"main" bson:"main"`
	Light      string `json:"light" bson:"light"`
	Dark       string `json:"dark" bson:"dark"`
	Error      string `json:"error" bson:"error"`
	Background string `json:"background" bson:"background"`
}

// PortalPageProfile provide page information
type PortalPageProfile struct {
	Id string `json:"id" bson:"_id"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// URL generated id
	IdURL string `json:"url_id" bson:"url_id"`

	// title of webpage
	Title string `json:"title" bson:"title"`

	// footer text
	Footer string `json:"footer" bson:"footer"`

	// interface features
	Interface struct {
		// Favicon    string `json:"favicon" bson:"-"`
		Logo       string `json:"logo" bson:"-"`
		LogoFooter string `json:"logo_footer" bson:"-"`
		Background string `json:"background" bson:"-"`

		// FaviconURL    string `json:"favicon_url" bson:"favicon_url"`
		LogoURL       string `json:"logo_url" bson:"logo_url"`
		LogoFooterURL string `json:"logo_footer_url" bson:"logo_footer_url"`
		BackgroundURL string `json:"background_url" bson:"background_url"`
		ButtonColor   string `json:"button_color" bson:"button_color"`

		Color pageColor `json:"color" bson:"color"`

		// color themes (new 1.3)
		Theme      string `json:"theme" bson:"theme"`
		ThemeType  string `json:"theme_type" bson:"theme_type"`
		ThemeColor string `json:"theme_color" bson:"theme_color"`
	} `json:"interface" bson:"interface"`

	// locales
	Locales map[Locale]string `json:"locales" bson:"locales"`

	// service aggrement per locale
	Agreements map[Locale]string `json:"agreements" bson:"agreements"`

	// service aggrement plain text (no locale)
	Agreement string `json:"agreement" bson:"agreement"`

	// support field (email / phone / etc) (no locale)
	Support string `json:"support" bson:"support"`
}

// PortalPageSchedule for schedule of page_id to be on loc_id URL
type PortalPageSchedule struct {
	ID       string   `json:"id" bson:"_id"`
	PageID   string   `json:"page_id" bson:"page_id"`
	PageURL  string   `json:"page_url" bson:"page_url"`
	LocID    string   `json:"loc_id" bson:"loc_id"`
	Schedule Schedule `json:"schedule" bson:"schedule"`
}

type AdImageObject struct {
	Text  string `json:"text" bson:"text"`
	Image string `json:"image" bson:"image"`
}

type PortalAdData struct {
	Type PortalAdvertisementType `json:"type" bson:"type"`

	// URL of iframe or URL to file on portal
	URL  string `json:"url" bson:"url"`
	File string `json:"file,omitempty" bson:"-"`

	// Desktop URL
	URLDesktop  string `json:"url_desktop" bson:"url_desktop"`
	FileDesktop string `json:"file_desktop,omitempty" bson:"-"`

	// question and poll variants if type == poll
	Question     string          `json:"question" bson:"question"`
	QuestionDesc string          `json:"question_desc" bson:"question_desc"`
	PollVariants []string        `json:"poll_variants" bson:"poll_variants"`
	PollImage    []AdImageObject `json:"poll_image" bson:"poll_image"`

	Duration     int64 `json:"duration" bson:"duration"`
	SkipDuration int64 `json:"skip_after" bson:"skip_after"`

	// for redirect to client's site
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`

	// title text
	Text string `json:"text" bson:"text"`

	// color options
	Color pageColor `json:"color,omitempty" bson:"color"`

	// color themes
	ThemeType  string `json:"theme_type" bson:"theme_type"`
	ThemeColor string `json:"theme_color" bson:"theme_color"`

	// поля для цвета кнопки и текста кнопки
	RedirectButtonColor string `json:"redirect_button_color" bson:"redirect_button_color"`
	RedirectButtonText  string `json:"redirect_button_text" bson:"redirect_button_text"`

	SelfVariant bool `json:"self_variant" bson:"self_variant"`
	Skip        bool `json:"skip" bson:"skip"`

	// поля означаюзие включен ли опрос данных о пользователя
	PollFields PollEnabledFields `json:"poll_fields" bson:"poll_fields"`
}

// PortalAd object for ihot
type PortalAd struct {
	Id string `json:"id" bson:"_id"`

	// common data
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// more specific data to link with profiles and authorization
	Profile       string `json:"profile" bson:"profile"`
	Authorization string `json:"authorization" bson:"authorization"`

	// number beetween 1 and 100 - more is higher
	Priority int `json:"priority" bson:"priority"`

	//OS operation system of client
	//OS PortalOSType `json:"os" bson:"os"`
	OS []string `json:"os" bson:"os"`

	//Vendor customer device manufacturers
	Vendor []string `json:"vendor" bson:"vendor"`

	//Platform descktop/mobile
	Platform struct {
		Desktop bool `json:"desktop" bson:"desktop"`
		Mobile  bool `json:"mobile" bson:"mobile"`
		Tablet  bool `json:"tablet" bson:"tablet"`
	} `json:"platform" bson:"platform"`

	// schedule of ads to work start-stop and number of views to show
	Schedule Schedule `json:"schedule" bson:"schedule"`

	NumberOfVisits struct {
		Visits int            `json:"visits" bson:"visits"`
		Sign   PortalSignType `json:"sign" bson:"sign"`
	} `json:"number_of_visits" bson:"number_of_visits"`

	UserRotation struct {
		Sex         string `json:"sex" bson:"sex"`
		DateOfBirth struct {
			Day   int `json:"day" bson:"day"`
			Month int `json:"month" bson:"month"`
			Year  int `json:"year" bson:"year"`
		} `json:"date_of_birth" bson:"date_of_birth"`
		Age  int64  `json:"age" bson:"age"`
		City string `json:"city" bson:"city"`
	} `json:"user_rotation" bson:"user_rotation"`

	LocIDs []string `json:"loc_ids" bson:"loc_ids"`

	Data PortalAdData `json:"data" bson:"data"`
}

type PortalAdStatRequest struct {
	Id            string `json:"id"`
	Profile       string `json:"profile"`
	LocID         string `json:"loc_id"`
	Authorization string `json:"authorization"`

	Duration          int64                    `json:"duration"`
	Skipped           bool                     `json:"skipped"`
	PollVariant       string                   `json:"poll_variant"`
	PollUserData      PortalPollUserDataAnswer `json:"poll_user_data_answer"`
	FollowRedirectURL bool                     `json:"follow_redirect"`
}

type PortalAdStatInc struct {
	Counter        int   `json:"counter" bson:"counter"`
	CounterSkipped int   `json:"counter_skipped" bson:"counter_skipped"`
	Duration       int64 `json:"duration" bson:"duration"`
	Follows        int   `json:"follows" bson:"follows"`
}

// PortalAdStat
type PortalAdStat struct {
	// same ID as portalad
	Id string `json:"id" bson:"_id"`

	// more specific data to link with profiles and authorization
	// Profile       string `json:"profile" bson:"profile"`
	// Authorization string `json:"authorization" bson:"authorization"`

	// Data PortalAdData `json:"data" bson:"data"`

	Counter        int            `json:"counter" bson:"counter"`
	CounterSkipped int            `json:"counter_skipped" bson:"counter_skipped"`
	Duration       int64          `json:"duration" bson:"duration"`
	PollCounter    map[string]int `json:"poll_counter" bson:"poll_counter"`
	Follows        int            `json:"follows" bson:"follows"`
}

type PortalAdStatDaily struct {
	Id   string `json:"id" bson:"_id"`
	IdAd string `json:"id_ad" bson:"id_ad"`

	// more specific data to link with profiles and authorization
	Profile       string `json:"profile" bson:"profile"`
	LocID         string `json:"loc_id"`
	Authorization string `json:"authorization" bson:"authorization"`

	// dat of day
	Year  int `json:"year" bson:"year"`
	Month int `json:"month" bson:"month"`
	Day   int `json:"day" bson:"day"`

	Counter        int            `json:"counter" bson:"counter"`
	CounterSkipped int            `json:"counter_skipped" bson:"counter_skipped"`
	Duration       int64          `json:"duration" bson:"duration"`
	PollCounter    map[string]int `json:"poll_counter" bson:"poll_counter"`
	Follows        int            `json:"follows" bson:"follows"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

type PortalAdStatLog struct {
	ID      string                  `json:"id" bson:"_id"`
	IDAd    string                  `json:"id_ad" bson:"id_ad"`
	Profile string                  `json:"profile" bson:"profile"`
	LocID   string                  `json:"loc_id" bson:"loc_id"`
	Account string                  `json:"account" bson:"account"`
	TypeAd  PortalAdvertisementType `json:"type_ad" bson:"type_ad"`

	CreateAt int64 `json:"create_at" bson:"create_at"`

	Time struct {
		Day   int `json:"day" bson:"day"`
		Month int `json:"month" bson:"month"`
		Year  int `json:"year" bson:"year"`
	} `json:"time" bson:"time"`

	UA UserAgent `json:"ua" bson:"ua"`

	UserAge         int    `json:"user_age" bson:"user_age"`
	UserDateOfBirth int64  `json:"user_date_of_birth" bson:"user_date_of_birth"`
	UserSex         string `json:"user_sex" bson:"user_sex"`

	CPE         string `json:"cpe" bson:"cpe"`
	MAC         string `json:"mac" bson:"mac"`
	WLAN        string `json:"wlan" bson:"wlan"`
	NASID       string `json:"nas_id" bson:"nas_id"`
	AdFollowURL string `json:"ad_follow_url" bson:"ad_follow_url"`
}

//PortalAuthorizationWithSN авторизация с помощью социальных сетей
type PortalAuthorizationWithSN struct {
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`
	GroupURL    string `json:"group_url" bson:"group_url"`
	Post        struct {
		Message     string `json:"message" bson:"message"`
		AttachedURL string `json:"attached_url" bson:"attached_url"`
	} `json:"post" bson:"post"`
}

//PortalPollUserDataAnswer for answer about user data(sex,date_of_birth etc)
type PortalPollUserDataAnswer struct {
	Sex         string `json:"sex" bson:"sex"`
	DateOfBirth struct {
		Day   int `json:"day" bson:"day"`
		Month int `json:"month" bson:"month"`
		Year  int `json:"year" bson:"year"`
	} `json:"date_of_birth" bson:"date_of_birth"`
}

//PollEnabledFields включение опросов о данных пользователя
type PollEnabledFields struct {
	DateOfBirth bool `json:"date_of_birth" bson:"date_of_birth"`
	Sex         bool `json:"sex" bson:"sex"`
}
