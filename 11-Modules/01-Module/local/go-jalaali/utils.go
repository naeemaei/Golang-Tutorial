package jalaali

import "strings"

var enToFa = strings.NewReplacer(
	"0", "۰",
	"1", "۱",
	"2", "۲",
	"3", "۳",
	"4", "۴",
	"5", "۵",
	"6", "۶",
	"7", "۷",
	"8", "۸",
	"9", "۹",
)

// IsValidDate take Jalaali date and return true if it is valid,
// otherwise false.
func IsValidDate(jy, jm, jd int) bool {
	d, err := MonthLength(jy, jm)
	if err != nil {
		return false
	}
	return -61 <= jy && jy <= 3177 &&
		1 <= jm && jm <= 12 &&
		1 <= jd && jd <= d
}

// MonthLength take Jalaali date and return length of that specific
// month. Error is not nil if Jalaali year passed to function is not valid.
func MonthLength(jy, jm int) (int, error) {
	if jm <= 6 {
		return 31, nil
	} else if jm <= 11 {
		return 30, nil
	}

	leap, err := IsLeapYear(jy)
	if err != nil {
		return 0, err
	} else if leap {
		return 30, nil
	}
	return 29, nil
}

// IsLeapYear take a Jalaali year and return true if it is leap year. Error
// is not nil if Jalaali year passed to function is not valid.
func IsLeapYear(jy int) (bool, error) {
	leap, _, _, err := jalCal(jy)
	return leap == 0, err
}