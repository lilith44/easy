package easy

import (
	"errors"
	"time"
)

type Time time.Time

func NewTime(t time.Time) Time {
	return Time(t.Local())
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	tm, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
	if err != nil {
		return err
	}

	*t = NewTime(tm)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	tm := t.Time()
	if y := tm.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = tm.AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

func (t Time) MarshalText() ([]byte, error) {
	tm := t.Time()
	if y := tm.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(time.DateTime))
	return tm.AppendFormat(b, time.DateTime), nil
}

func (t *Time) UnmarshalText(data []byte) error {
	tm, err := time.ParseInLocation(time.DateTime, string(data), time.Local)
	if err != nil {
		return err
	}

	*t = NewTime(tm)
	return nil
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t Time) String() string {
	return t.Time().Format(time.DateTime)
}

func Now() Time {
	return NewTime(time.Now())
}
