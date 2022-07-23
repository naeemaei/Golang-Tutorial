package jalaali

import (
	"testing"
	"time"
)

func TestFromYMD(t *testing.T) {
	tests := []struct {
		gy, gm, gd, jy, jm, jd int
	}{
		{1981, 8, 17, 1360, 5, 26},
		{2013, 1, 10, 1391, 10, 21},
		{2014, 8, 4, 1393, 5, 13},
	}

	for _, test := range tests {
		y, m, d, err := ToJalaali(test.gy, time.Month(test.gm), test.gd)
		if err != nil {
			t.Errorf("%v", err)
		} else if y != test.jy || m != Month(test.jm) || d != test.jd {
			t.Errorf("Expected %v/%v/%v got %v/%v%v.", test.jy, test.jm, test.jd, y, m, d)
		}
	}

}

func TestToGregorian(t *testing.T) {
	tests := []struct {
		jy, jm, jd, gy, gm, gd int
	}{
		{1360, 5, 26, 1981, 8, 17},
		{1391, 10, 21, 2013, 1, 10},
		{1393, 5, 13, 2014, 8, 4},
	}

	for _, test := range tests {
		y, m, d, err := ToGregorian(test.jy, Month(test.jm), test.jd)
		if err != nil {
			t.Errorf("%v", err)
		} else if y != test.gy || m != time.Month(test.gm) || d != test.gd {
			t.Errorf("Expected %v/%v/%v got %v/%v%v.", test.gy, test.gm, test.gd, y, m, d)
		}
	}

}

func TestIsValidDate(t *testing.T) {
	tests := []struct {
		y, m, d int
		ok      bool
	}{
		{-62, 12, 29, false},
		{-61, 1, 1, true},
		{3178, 1, 1, false},
		{3177, 12, 29, true},
		{1393, 0, 1, false},
		{1393, 13, 1, false},
		{1393, 1, 0, false},
		{1393, 1, 32, false},
		{1393, 1, 31, true},
		{1393, 11, 31, false},
		{1393, 11, 30, true},
		{1393, 12, 30, false},
		{1393, 12, 29, true},
		{1395, 12, 30, true},
	}

	for _, test := range tests {
		valid := IsValidDate(test.y, test.m, test.d)
		if valid != test.ok {
			calculated, actual := "", " not"
			if test.ok {
				calculated, actual = " not", ""
			}
			t.Errorf("%v/%v/%v is%v valid date but considered%v valid.",
				test.y, test.m, test.d, actual, calculated)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year int
		leap bool
	}{
		{1393, false},
		{1394, false},
		{1395, true},
		{1396, false},
	}

	for _, test := range tests {
		leap, err := IsLeapYear(test.year)
		if err != nil {
			t.Errorf("%v", err)
		} else if leap != test.leap {
			calculated, actual := "", " not"
			if leap {
				calculated, actual = " not", ""
			}
			t.Errorf("%v is%v leap but considered%v leap.", test.year, actual, calculated)
		}
	}
}

func TestMonthLength(t *testing.T) {
	tests := []struct {
		y, m, ml int
	}{
		{1393, 1, 31},
		{1393, 4, 31},
		{1393, 6, 31},
		{1393, 7, 30},
		{1393, 10, 30},
		{1393, 12, 29},
		{1394, 12, 29},
		{1395, 12, 30},
	}

	for _, test := range tests {
		calculated, err := MonthLength(test.y, test.m)
		if err != nil {
			t.Errorf("%v", err)
		} else if calculated != test.ml {
			t.Errorf("Length of %v/%v month is %v but considered %v.",
				test.y, test.m, test.ml, calculated)
		}
	}
}

func TestJFormat(t *testing.T) {
	iran, _ := time.LoadLocation("Asia/Tehran")

	tests := []struct {
		time   time.Time
		format []string
		result []string
	}{
		{
			time.Date(2001, 1, 1, 1, 1, 1, 1, iran),
			[]string{
				"2006 06", // Year formatting
				"January Jan 1 01", // Month formatting
				"Monday Mon 2 _2 02", // Day formatting
				"15 3 03 4 04 5 05 PM pm", // Hour, Minute, Second formatting
				".0 .00 .000 .000000 .000000000 .9 .99 .999 .999999 .999999999", // Nanosecond formatting
			},
			[]string{
				"۱۳۷۹ ۷۹", // Year formatting
				"دی دی ۱۰ ۱۰", // Month formatting
				"دوشنبه دوشنبه ۱۲ ۱۲ ۱۲", // Day formatting
				"۰۱ ۱ ۰۱ ۱ ۰۱ ۱ ۰۱ قبل‌ازظهر قبل‌ازظهر", // Hour, Minute, Second formatting
				".۰ .۰۰ .۰۰۰ .۰۰۰۰۰۰ .۰۰۰۰۰۰۰۰۱     .۰۰۰۰۰۰۰۰۱", // Nanosecond formatting
			},
		}, {
			time.Date(2001, 2, 3, 15, 17, 1, 999999999, iran),
			[]string{
				"2006 06", // Year formatting
				"January Jan 1 01", // Month formatting
				"Monday Mon 2 _2 02", // Day formatting
				"15 3 03 4 04 5 05 PM pm", // Hour, Minute, Second formatting
				".0 .00 .000 .000000 .000000000 .9 .99 .999 .999999 .999999999", // Nanosecond formatting
			},
			[]string{
				"۱۳۷۹ ۷۹", // Year formatting
				"بهمن بهمن ۱۱ ۱۱", // Month formatting
				"شنبه شنبه ۱۵ ۱۵ ۱۵", // Day formatting
				"۱۵ ۳ ۰۳ ۱۷ ۱۷ ۱ ۰۱ بعدازظهر بعدازظهر", // Hour, Minute, Second formatting
				".۹ .۹۹ .۹۹۹ .۹۹۹۹۹۹ .۹۹۹۹۹۹۹۹۹ .۹ .۹۹ .۹۹۹ .۹۹۹۹۹۹ .۹۹۹۹۹۹۹۹۹", // Nanosecond formatting
			},
		},
	}

	for i, test := range tests {
		j := From(test.time)

		for f := range test.format {
			result, err := j.JFormat(test.format[f])
			if err != nil {
				t.Error(err)
			}
			if result != test.result[f] {
				t.Error("Bad formatting for test as index: ", i, "\nWanted: ", test.result[f], "\nGot:    ", result)
			}
		}
	}
}
