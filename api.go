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
