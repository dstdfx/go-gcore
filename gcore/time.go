package gcore

import (
	"strings"
	"time"
)

const DateFormat = "2006-01-02T15:04:05"

// GCoreTime represents custom time type
type GCoreTime struct {
	time.Time
}

func NewGCoreTime(t time.Time) *GCoreTime {
	return &GCoreTime{Time: t}
}

func (t *GCoreTime) UnmarshalJSON(b []byte) (err error) {
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
