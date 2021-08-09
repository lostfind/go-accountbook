package model

import (
	"testing"
	"time"
)

func TestSetYearMonth(t *testing.T) {
	testCases := []struct {
		ym        string
		want      time.Time
		wantError string
	}{
		{"2020-04", time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local), ""},
		{"202005", *new(time.Time), "Parse err"},
	}

	for i, tc := range testCases {
		budget := new(Budget)

		err := budget.SetYearMonth(tc.ym)
		if err != nil && err.Error() != tc.wantError {
			t.Errorf("%d) Error is %s; want %s", i+1, err, tc.wantError)
		}

		if got := budget.YearMonth; got != tc.want {
			t.Errorf("%d) SetYearMonth(%s) = %s; want %s", i+1, tc.ym, got, tc.want)
		}
	}
}
