package libportal

// APIRequest struct for common HTTP
// request from other services (like backend)
type APIRequest struct {

	// ids for onject query
	Ids []string `json:"ids" query:"ids[]"`

	// Pagination
	Sort   string `json:"sort" query:"sort"`
	Offset int    `json:"offset" query:"offset"`
	Limit  int    `json:"limit" query:"limit"`

	// for Search
	SearchQuery string `json:"search,omitempty" query:"search"`

	// profiles ids for many objects and access control
	Profiles []string `json:"profiles,omitempty" query:"profiles[]"`

	// account ids for accounts data if needed
	Accounts []string `json:"accounts,omitempty" query:"accounts[]"`

	// tariffs ids for vouchers if needed
	Tariffs []string `json:"tariffs,omitempty" query:"tariffs[]"`

	//AdType  type of ad to download
	AdTypes []string `json:"ad_types" query:"ad_types[]"`

	// account for black/white list
	ActionList    PortalActionListType `json:"action_list" query:"action_list"`
	AccountToList string               `json:"account_to_list" query:"account_to_list"`

	// query for stat (range / pie / timeseries)
	Start   int64  `json:"start" query:"start"`
	Stop    int64  `json:"stop" query:"stop"`
	Object  string `json:"object" query:"object"`
	Subtype string `json:"subtype" query:"subtype"`

	//Query or filter account by DateOfBirth
	Day   int `json:"day" query:"day"`
	Month int `json:"month" query:"month"`
	Year  int `json:"year" query:"year"`

	LengthVoucher int `json:"length_voucher" query:"length_voucher"`
	// if true, get return list code and start stop(about vouchers)
	ListVoucher bool `json:"list_voucher" query:"list_voucher"`

	NumberOfVouchers int `json:"number_of_vouchers" query:"number_of_vouchers"`
	// for transaction read -- possible values ["day", "week", "month", "quarter"]
	Last string `json:"last,omitempty" query:"last"`

	ClonePage bool `json:"clone_page" query:"clone_page"`

	MACMask bool `json:"mac_mask" query:"mac_mask"`

	Query string `json:"q" query:"q"`

	Request []map[string]interface{} `json:"request" query:"request" form:"request"`
}

// APIResponse struct for common response
type APIResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`

	Items  interface{} `json:"data,omitempty"`
	Total  int         `json:"total,omitempty"`
	Errors []string    `json:"errors,omitempty"`
}

// PortalBackendRequest struct for request
// from frontend to backend
type PortalBackendRequest struct {
	// Needed client data
	MAC     string `json:"mac" bson:"mac" form:"mac" query:"mac"`
	CPE     string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id"`
	WLAN    string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id"`
	Ip      string `json:"client_ip" bson:"client_ip" form:"client_ip" query:"client_ip"`
	NAS     string `json:"nas" bson:"nas"`
	GroupID string `json:"group_id" bson:"group_id"`

	// client credentials
	Username string `json:"username,omitempty" bson:"username" form:"username" query:"username" validate:"-"`
	Password string `json:"password,omitempty" bson:"password" form:"password" query:"password" validate:"-"`

	// browser specific data
	UserAgent string `json:"useragent"  bson:"useragent" form:"useragent" query:"useragent" validate:"-"`

	// Address of platform CoA manager
	SwitchURL string `json:"switch_url" validate:"-"`

	// Remember period for user accounts
	Remember int64 `json:"remember"`

	// push aggrement
	PushAgreement bool `json:"push_agree"`

	// Type of choosen type
	Type string `json:"type"`

	// info for watched advertisement / poll
	Ad PortalAdStatRequest `json:"ad"`

	// auth type -- free / sponsor / paid / etc
	AuthType string `json:"auth_type"`

	// authentype -- sms / callback / etc
	AuthenType string `json:"authen_type"`

	// voucher for paid voucher internet
	Voucher string `json:"voucher"`

	// account update data
	AccountName    string `json:"account_name"`
	AccountSurName string `json:"account_surname"`

	// tarriff to buy a voucher
	Tariff string `json:"tariff"`

	// payment system
	PaymentSystem string `json:"payment_system"`
	PaymentAmount int    `json:"payment_amount"`

	// vouchers
	Vouchers []string `json:"vouchers"`

	// client locale
	Locale string `json:"locale"`

	// for voucher activation in once
	Activate bool `json:"activate"`

	//socialNetwork data
	SocialNetwork map[string]AccountFromSocialNetwork `json:"social_network" bson:"social_network"`

	// for internal using
	Timeout int64 `json:"-" validate:"-"`
	// profile id
	Profile string `json:"-" validate:"-"`
}

// PortalResponseObject struct for answer from Portal
// backend to frontend
type PortalBackendResponse struct {
	// status - success / error
	Status PortalResponseStatus `json:"status,omitempty"`
	// code 0 if OK, > 1 another
	Code int `json:"code"`
	// desctiption if error
	Description string `json:"description,omitempty"`
	// current state - authen, auth, ad, pass
	State PortalUserState `json:"state,omitempty"`
	// current substate - need, check
	Substate string `json:"substate,omitempty"`
	// additional data if needed
	Data interface{} `json:"data,omitempty"`

	// additional available data
	Available interface{} `json:"available,omitempty"`
	// account data provided
	Account interface{} `json:"account,omitempty"`

	// data for new auth stage
	Data2 interface{} `json:"data2,omitempty"`
}
