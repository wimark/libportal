package libportal

// GuestControlSettings for WLAN/Wired Guest Control settings
type GuestControlSettings struct {
	CaptiveRedirect string   `json:"captive_redirect" bson:"captive_redirect"`
	MACAuth         []string `json:"mac_radius_auth_servers" bson:"mac_radius_auth_servers"`
}

// DnsAddress for link IP with domain name
type DnsAddress struct {
	Ip         string `json:"ip" bson:"ip"`
	DomainName string `json:"domain_name" bson:"domain_name"`
}

// CaptiveRedirect struct for manage of Client redirect system
type CaptiveRedirect struct {
	Name string `json:"name" bson:"name"`

	// URL to redirect (Portal Web part)
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`

	MACWhiteList []string `json:"mac_list" bson:"mac_list"`

	URLWhiteList []DnsAddress `json:"url_list" bson:"url_list"`
	PreAuthList  []string     `json:"preauth_list" bson:"preauth_list"`

	NoMasquerade bool `json:"no_masquerade" bson:"no_masquerade"`
}
