package types

import (
	"encoding/json"
	"time"
)

const DateFormat = "2006-01-02"

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(DateFormat))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

func (d Date) String() string {
	return d.Time.Format(DateFormat)
}

func (d *Date) UnmarshalText(data []byte) error {
	parsed, err := time.Parse(DateFormat, string(data))
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var timeStr string
	err := json.Unmarshal(data, &timeStr)
	if err != nil {
		return err
	}
	parsed, err1 := time.Parse(time.RFC3339, timeStr)
	if err1 == nil {
		t.Time = parsed
		return nil
	}

	parsed2, err2 := time.Parse("2006-01-02 15:04:05-07:00", timeStr)
	if err2 == nil {
		t.Time = parsed2
		return nil
	}

	return err1
}
