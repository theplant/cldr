package cldr_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	"github.com/razor-1/cldr/resources"
)

const dateTimeString = "Jan 2, 2006 at 3:04:05pm"

func TestFormatDateTime(t *testing.T) {
	datetime, _ := time.Parse(dateTimeString, dateTimeString)
	ta := assert.New(t)
	en, err := resources.GetLocale(language.Make("en"))
	ta.NoError(err)

	tests := []struct {
		inTime  time.Time
		err     error
		fmtFunc func(time.Time) (string, error)
		format  string
	}{
		{fmtFunc: en.Calendar.FmtDateFull, format: "Monday, January 2, 2006"},
		{fmtFunc: en.Calendar.FmtDateLong, format: "January 2, 2006"},
		{fmtFunc: en.Calendar.FmtDateMedium, format: "Jan 2, 2006"},
		{fmtFunc: en.Calendar.FmtDateShort, format: "1/2/06"},
		{fmtFunc: en.Calendar.FmtTimeFull, format: "3:04:05 PM GMT+00:00"},
		{fmtFunc: en.Calendar.FmtTimeLong, format: "3:04:05 PM GMT+00:00"},
		{fmtFunc: en.Calendar.FmtTimeMedium, format: "3:04:05 PM"},
		{fmtFunc: en.Calendar.FmtTimeShort, format: "3:04 PM"},
		{fmtFunc: en.Calendar.FmtDateTimeFull, format: "Monday, January 2, 2006 at 3:04:05 PM GMT+00:00"},
		{fmtFunc: en.Calendar.FmtDateTimeLong, format: "January 2, 2006 at 3:04:05 PM GMT+00:00"},
		{fmtFunc: en.Calendar.FmtDateTimeMedium, format: "Jan 2, 2006, 3:04:05 PM"},
		{fmtFunc: en.Calendar.FmtDateTimeShort, format: "1/2/06, 3:04 PM"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			var tm time.Time
			if test.inTime.IsZero() {
				// all tests use datetime. default to it but still allow an easy override for future tests
				tm = datetime
			} else {
				tm = test.inTime
			}
			out, err := test.fmtFunc(tm)
			ta.Equal(test.err, err)
			ta.Equal(test.format, out)
		})
	}

	r, err := en.Calendar.Format(datetime, "MMMM d yy")
	ta.NoError(err)
	ta.Equal("January 2 06", r)
}
