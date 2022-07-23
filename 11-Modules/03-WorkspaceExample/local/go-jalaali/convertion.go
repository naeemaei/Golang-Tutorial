package jalaali

import (
	"errors"
	"time"
)

var (
	breaks = [...]int{-61, 9, 38, 199, 426, 686, 756, 818, 1111, 1181, 1210,
		1635, 2060, 2097, 2192, 2262, 2324, 2394, 2456, 3178}
)

// ToJalaali converts Gregorian to Jalaali date. Error is not nil if Jalaali
// year passed to function is not valid.
func ToJalaali(gregorianYear int, gregorianMonth time.Month, gregorianDay int) (int, Month, int, error) {
	jy, jm, jd, err := d2j(g2d(gregorianYear, int(gregorianMonth), gregorianDay))
	return jy, Month(jm), jd, err
}

// ToGregorian converts Jalaali to Gregorian date. Error is not nil if Jalaali
// year passed to function is not valid.
func ToGregorian(jalaaliYear int, jalaaliMonth Month, jalaaliDay int) (int, time.Month, int, error) {

	if jalaaliMonth > 12 || jalaaliMonth < 1 {
		return 0, 0, 0, errors.New("Invalid Jalaali month")
	}

	jdn, err := j2d(jalaaliYear, int(jalaaliMonth), jalaaliDay)
	if err != nil {
		return 0, 0, 0, err
	}

	gy, gm, gd := d2g(jdn)
	return gy, time.Month(gm), gd, nil
}

func j2d(jy, jm, jd int) (jdn int, err error) {
	_, gy, march, err := jalCal(jy)
	if err != nil {
		return 0, err
	}
	return g2d(gy, 3, march) + (jm-1)*31 - div(jm, 7)*(jm-7) + jd - 1, nil
}

func d2j(jdn int) (int, int, int, error) {
	gy, _, _ := d2g(jdn) // Calculate Gregorian year (gy).
	jy := gy - 621
	leap, _, march, err := jalCal(jy)
	jdn1f := g2d(gy, 3, march)

	if err != nil {
		return 0, 0, 0, err
	}

	// Find number of days that passed since 1 Farvardin.
	k := jdn - jdn1f
	if k >= 0 {
		if k <= 185 {
			// The first 6 months.
			jm := 1 + div(k, 31)
			jd := mod(k, 31) + 1
			return jy, jm, jd, nil
		}
		// The remaining months.
		k -= 186
	} else {
		// Previous Jalaali year.
		jy--
		k += 179
		if leap == 1 {
			k++
		}
	}
	jm := 7 + div(k, 30)
	jd := mod(k, 30) + 1
	return jy, jm, jd, nil
}

func jalCal(jy int) (int, int, int, error) {
	bl, gy, leapJ, jp := len(breaks), jy+621, -14, breaks[0]
	jump := 0

	if jy < jp || jy >= breaks[bl-1] {
		return 0, 0, 0, &ErrorInvalidYear{jy}
	}

	// Find the limiting years for the Jalaali year jy.
	for i := 1; i < bl; i++ {
		jm := breaks[i]
		jump = jm - jp
		if jy < jm {
			break
		}
		leapJ += div(jump, 33)*8 + div(mod(jump, 33), 4)
		jp = jm
	}
	n := jy - jp

	// Find the number of leap years from AD 621 to the beginning
	// of the current Jalaali year in the Persian calendar.
	leapJ += div(n, 33)*8 + div(mod(n, 33)+3, 4)
	if mod(jump, 33) == 4 && jump-n == 4 {
		leapJ++
	}

	// And the same in the Gregorian calendar (until the year gy).
	leapG := div(gy, 4) - div((div(gy, 100)+1)*3, 4) - 150

	// Determine the Gregorian date of Farvardin the 1st.
	march := 20 + leapJ - leapG

	// Find how many years have passed since the last leap year.
	if jump-n < 6 {
		n -= jump + div(jump+4, 33)*33
	}
	leap := mod(mod(n+1, 33)-1, 4)
	if leap == -1 {
		leap = 4
	}

	return leap, gy, march, nil
}

func g2d(gy, gm, gd int) int {
	d := div((gy+div(gm-8, 6)+100100)*1461, 4) +
		div(153*mod(gm+9, 12)+2, 5) +
		gd - 34840408
	d = d - div(div(gy+100100+div(gm-8, 6), 100)*3, 4) + 752
	return d
}

func d2g(jdn int) (int, int, int) {
	j := 4*jdn + 139361631
	j = j + div(div(4*jdn+183187720, 146097)*3, 4)*4 - 3908
	i := div(mod(j, 1461), 4)*5 + 308
	gd := div(mod(i, 153), 5) + 1
	gm := mod(div(i, 153), 12) + 1
	gy := div(j, 1461) - 100100 + div(8-gm, 6)
	return gy, gm, gd
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}
