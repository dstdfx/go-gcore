package gcore

import (
	"strings"
	"time"
)

const DateFormat = "2006-01-02T15:04:05"

// Time represents custom time type
type Time struct {
	time.Time
}

// NewTime returns new instance of Time.
func NewTime(t time.Time) *Time {
	return &Time{Time: t}
}

// UnmarshalJSON represents custom implementation of Unmarshaler interface.
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	if strings.HasSuffix(s, "Z") {
		s = s[:len(s)-1]
	}

	t.Time, err = time.Parse(DateFormat, s)
	return
}
