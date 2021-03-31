package libportal

import "time"

const (
	CollPortalAccounting = "portal_accounting"
)

type RadAccessRequest struct {
	NASIdentifier string `json:"NAS-Identifier"`
	NASIPAddress  string `json:"NAS-IP-Address"`
	NASPortId     string `json:"NAS-Port-Id"`
	NASPort       string `json:"NAS-Port"`
	NASPortType   string `json:"NAS-Port-Type"`

	UserName         string `json:"User-Name"`
	CHAPPassword     string `json:"CHAP-Password"`
	CHAPChallenge    string `json:"CHAP-Challenge"`
	FramedIPAddress  string `json:"Framed-IP-Address"`
	WISPrLocationID  string `json:"WISPr-Location-ID"`
	WISPrLogoffURL   string `json:"WISPr-Logoff-URL"`
	CalledStationId  string `json:"Called-Station-Id"`
	CallingStationId string `json:"Calling-Station-Id"`
	AcctSessionId    string `json:"Acct-Session-Id"`

	ServiceType          string `json:"Service-Type"`
	MessageAuthenticator string `json:"Message-Authenticator"`
	EventTimestamp       string `json:"Event-Timestamp"`
	ChilliSpotVersion    string `json:"ChilliSpot-Version"`
}

type RadAccessResponce struct {
	Code  int         `json:"Code"`
	Reply []mapEntity `json:"Reply-Message"`
}

type RadAcctRequest struct {
	CalledStationId  string `json:"Called-Station-Id"`
	CallingStationId string `json:"Calling-Station-Id"`
	EventTimestamp   string `json:"Event-Timestamp"`

	NASIdentifier string `json:"NAS-Identifier"`
	NASIPAddress  string `json:"NAS-IP-Address"`
	NASPortType   string `json:"NAS-Port-Type"`

	AcctStatusType      string `json:"Acct-Status-Type"`
	AcctTerminateCause  string `json:"Acct-Terminate-Cause,omitempty"`
	AcctSessionID       string `json:"Acct-Session-ID`
	AcctDelayTime       int    `json:"Acct-Delay-Time"`
	AcctInputGigawords  int    `json:"Acct-Input-Gigawords"`
	AcctOutputGigawords int    `json:"Acct-Output-Gigawords"`
	AcctOutputOctets    int    `json:"Acct-Output-Octets"`
	AcctInputOctets     int    `json:"Acct-Input-Octets"`
	AcctInputPackets    int    `json:"Acct-Input-Packets"`
	AcctOutputPackets   int    `json:"Acct-Output-Packets"`
	AcctSessionTime     int    `json:"Acct-Session-Time"`

	Attr261455910     string `json:"Attr-26.14559.10"`
	ChilliSpotVersion string `json:"ChilliSpot-Version"`
	WISPrLocationID   string `json:"WISPr-Location-ID"`
}

type RadAcctResponse RadAccessResponce

type RadAccounting struct {
	RadAcctRequest
	Create time.Time `json:"create" bson:"create"`
	ID     string    `json:"id" bson:"_id"`
}

type mapEntity struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}
