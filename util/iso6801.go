package util

import (
	"fmt"
	"time"
)

//Get timestamp string with ISO8601 format
func GetISO8601TimeStamp(ts time.Time) string {
	t := ts.UTC()
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

const format_ISO8601 = "2006-01-02T15:04:05Z"
const jsonFormat_ISO8601 = `"` + format_ISO8601 + `"`

type ISO6801Time time.Time

// New constructs a new iso8601.Time instance from an existing
// time.Time instance.  This causes the nanosecond field to be set to
// 0, and its time zone set to a fixed zone with no offset from UTC
// (but it is *not* UTC itself).
func New(t time.Time) ISO6801Time {
	return ISO6801Time(time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.UTC,
	))
}

func (it *ISO6801Time) IsDefault() bool {
	return *it == ISO6801Time{}
}

func (it ISO6801Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(it).Format(jsonFormat_ISO8601)), nil
}

func (it *ISO6801Time) UnmarshalJSON(data []byte) error {
	if string(data) == "\"\"" {
		return nil
	}
	t, err := time.ParseInLocation(jsonFormat_ISO8601, string(data), time.UTC)
	if err == nil {
		*it = ISO6801Time(t)
	}
	return err
}

func (it ISO6801Time) String() string {
	return time.Time(it).String()
}