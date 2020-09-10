package libportal

import (
	"time"
)

// Schedule
type Schedule struct {
	// only for ADvertisement
	Views int `json:"views" bson:"views"`

	// range of schedule
	Start int64 `json:"start" bson:"start"`
	Stop  int64 `json:"stop" bson:"stop"`

	// diff with UTC (+4 or -3)
	UTCDiff int `json:"utc_diff" bson:"utc_diff"`

	// day Of week mask (1 << ct.Weekday()!)
	// example: 1 << 0 && 1 << 6 for Sunday && Saturday (or yeah, Sunday FIRST)
	DoWMask int `json:"dow_mask" bson:"dow_mask"`

	// hour of day mask (1 << ct.Hour()!)
	// example: HourMask = 1 << 0 + 1 << 1 for 0-1 and 1-2 hours
	HourMask int `json:"hour_mask" bson:"hour_mask"`

	// old time of day format
	TimeOfDay struct {
		Morning   bool `json:"morning,omitempty" bson:"morning"`
		Afternoon bool `json:"afternoon,omitempty" bson:"afternoon"`
		Evening   bool `json:"evening,omitempty" bson:"evening"`
		Night     bool `json:"night,omitempty" bson:"night"`
	} `json:"time_of_day,omitempty" bson:"time_of_day"`

	// old day of week format
	DayOfWeek struct {
		Monday    bool `json:"monday,omitempty" bson:"monday"`
		Tuesday   bool `json:"tuesday,omitempty" bson:"tuesday"`
		Wednesday bool `json:"wednesday,omitempty" bson:"wednesday"`
		Thursday  bool `json:"thursday,omitempty" bson:"thursday"`
		Friday    bool `json:"friday,omitempty" bson:"friday"`
		Saturday  bool `json:"saturday,omitempty" bson:"saturday"`
		Sunday    bool `json:"sunday,omitempty" bson:"sunday"`
	} `json:"day_of_week,omitempty" bson:"day_of_week"`
}

// At method to check is ct at schedule or not
func (s *Schedule) At(ct int64) bool {
	if (s.Start > 0 && ct < s.Start) || ct > s.Stop {
		return false
	}
	t := time.Unix(ct+int64(s.UTCDiff*3600), 0)
	return s.atWeekday(int(t.Weekday())) && s.atHour(t.Hour())
}

// atWeekday to check is weekday by mask
func (s *Schedule) atWeekday(weekday int) bool {
	return s.DoWMask == 0 || (s.DoWMask&int(1<<(weekday%7))) > 0
}

// atHour to check is hour by mask
func (s *Schedule) atHour(hour int) bool {
	return s.HourMask == 0 || (s.HourMask&1<<(hour%24)) > 0
}
