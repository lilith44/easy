package easy

import (
	"errors"
	"time"
)

var DateTime = "2006-01-02 15:04:05"

type Time time.Time

func NewTime(tm time.Time) Time {
	return Time(tm)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	tm, err := time.Parse(`"`+DateTime+`"`, string(data))
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

	b := make([]byte, 0, len(DateTime)+2)
	b = append(b, '"')
	b = tm.AppendFormat(b, DateTime)
	b = append(b, '"')
	return b, nil
}

func (t Time) MarshalText() ([]byte, error) {
	tm := t.Time()
	if y := tm.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateTime))
	return tm.AppendFormat(b, DateTime), nil
}

func (t *Time) UnmarshalText(data []byte) error {
	tm, err := time.Parse(`"`+DateTime+`"`, string(data))
	if err != nil {
		return err
	}

	*t = NewTime(tm)
	return nil
}

func (t *Time) UnmarshalBSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	tm, err := time.Parse(`"`+DateTime+`"`, string(data))
	if err != nil {
		return err
	}

	*t = NewTime(tm)
	return nil
}

func (t Time) MarshalBSON() ([]byte, error) {
	tm := t.Time()
	if y := tm.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalBSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateTime)+2)
	b = append(b, '"')
	b = tm.AppendFormat(b, DateTime)
	b = append(b, '"')
	return b, nil
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t Time) String() string {
	return t.Time().Format(DateTime)
}

func Now() Time {
	return NewTime(time.Now())
}

func SetTimeLayout(layout string) {
	DateTime = layout
}
