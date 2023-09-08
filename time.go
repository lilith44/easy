package easy

import (
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

	tm, err := time.ParseInLocation(`"`+time.DateTime+`"`, ByteToString(data), time.Local)
	if err != nil {
		return err
	}

	*t = NewTime(tm)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	format := t.Time().Format(`"` + time.DateTime + `"`)
	return StringToByte(format), nil
}

func (t Time) MarshalText() ([]byte, error) {
	format := t.Time().Format(time.DateTime)
	return StringToByte(format), nil
}

func (t *Time) UnmarshalText(data []byte) error {
	tm, err := time.ParseInLocation(time.DateTime, ByteToString(data), time.Local)
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
