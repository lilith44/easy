package easy

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTime_MarshalJSON(t *testing.T) {
	format := "2023-09-01 00:00:00"
	want := `"2023-09-01 00:00:00"`
	t1, _ := time.ParseInLocation(time.DateTime, format, time.Local)
	t2 := Time(t1)

	got, err := json.Marshal(t2)
	if err != nil {
		t.Fatalf("error occurs while json.Marshal: %s", err)
	}

	if string(got) != want {
		t.Errorf("(%v).MarshalJSON() = %v, want %v", t2, string(got), want)
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	data := `"2023-09-01 00:00:00"`
	want := "2023-09-01 00:00:00"

	var tm Time
	if err := json.Unmarshal([]byte(data), &tm); err != nil {
		t.Fatalf("error occurs while json.Unmarshal: %s", err)
	}

	if tm.Time().Format(time.DateTime) != want {
		t.Errorf("(%v).UnmarshalJSON() = %v, want %v", data, tm, want)
	}
}
