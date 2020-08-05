package libportal

import (
	ua "github.com/mileusna/useragent"
)

const (
	uaMobile  = "mobile"
	uaTablet  = "tablet"
	uaDesktop = "desktop"
	uaBot     = "bot"
)

// UserAgent struct for parse useragent to internal struct
type UserAgent struct {
	Type      string `json:"type" bson:"type"`
	Device    string `json:"device" bson:"device"`
	Name      string `json:"name" bson:"name"`
	OS        string `json:"os" bson:"os"`
	OSVersion string `json:"os_version" bson:"os_version"`
}

func NewUserAgent(useragent string) UserAgent {

	var ua = ua.Parse(useragent)
	var u = UserAgent{
		OS:        ua.OS,
		OSVersion: ua.OSVersion,
		Device:    ua.Device,
		Name:      ua.Name,
	}

	var uaType string
	if ua.Mobile {
		uaType = uaMobile
	}
	if ua.Tablet {
		uaType = uaTablet
	}
	if ua.Desktop {
		uaType = uaDesktop
	}
	if ua.Bot {
		uaType = uaBot
	}
	u.Type = uaType
	return u
}
