package jalaali

const (
	_              = iota
	stdLongMonth   = iota + stdNeedDate  // "January"
	stdMonth                             // "Jan"
	stdNumMonth                          // "1"
	stdZeroMonth                         // "01"
	stdLongWeekDay                       // "Monday"
	stdWeekDay                           // "Mon"
	stdDay                               // "2"
	stdUnderDay                          // "_2"
	stdZeroDay                           // "02"
	stdHour        = iota + stdNeedClock // "15"
	stdHour12                            // "3"
	stdZeroHour12                        // "03"
	stdMinute                            // "4"
	stdZeroMinute                        // "04"
	stdSecond                            // "5"
	stdZeroSecond                        // "05"
	stdLongYear    = iota + stdNeedDate  // "2006"
	stdYear                              // "06"
	stdPM          = iota + stdNeedClock // "PM"
	stdpm                                // "pm"
	stdFracSecond0                       // ".0", ".00", ... , trailing zeros included
	stdFracSecond9                       // ".9", ".99", ..., trailing zeros omitted

	stdNeedDate  = 1 << 8             // need month, day, year
	stdNeedClock = 2 << 8             // need hour, minute, second
	stdArgShift  = 16                 // extra argument in high bits, above low stdArgShift
	stdMask      = 1<<stdArgShift - 1 // mask out argument
)

// std0x records the std values for "01", "02", ..., "06".
var std0x = [...]int{stdZeroMonth, stdZeroDay, stdZeroHour12, stdZeroMinute, stdZeroSecond, stdYear}

// JFormat gets default Golang layout string and parse put Jalaali calender information
// into the final string and return it.
func (j Jalaali) JFormat(layout string) (string, error) {
	const minBufSize = 64

	bufSize := len(layout)
	if bufSize < minBufSize { // minimum buffer size
		bufSize = minBufSize
	}
	b := make([]byte, 0, len(layout))

	b, err := j.jAppendFormat(b, layout)
	return string(b), err
}

// jAppendFormat is like JFormat but appends the textual
// representation to b and returns the extended buffer.
func (j Jalaali) jAppendFormat(b []byte, layout string) ([]byte, error) {
	var (
		year  int = -1
		month Month
		day   int
		hour  int = -1
		min   int
		sec   int
	)
	// Each iteration generates one std value.
	for layout != "" {
		prefix, std, suffix := nextStdChunk(layout)
		if prefix != "" {
			b = append(b, prefix...)
		}
		if std == 0 {
			break
		}
		layout = suffix

		// Compute year, month, day if needed.
		if year < 0 && std&stdNeedDate != 0 {
			var err error
			year, month, day, err = ToJalaali(j.Year(), j.Month(), j.Day())
			if err != nil {
				return b, err
			}
		}

		// Compute hour, minute, second if needed.
		if hour < 0 && std&stdNeedClock != 0 {
			hour, min, sec = j.Hour(), j.Minute(), j.Second()
		}

		switch std & stdMask {
		case stdYear:
			y := year
			if y < 0 {
				y = -y
			}
			b = appendInt(b, y%100, 2)
		case stdLongYear:
			b = appendInt(b, year, 4)
		case stdMonth, stdLongMonth:
			m := month.String()
			b = append(b, m...)
		case stdNumMonth:
			b = appendInt(b, int(month), 0)
		case stdZeroMonth:
			b = appendInt(b, int(month), 2)
		case stdWeekDay, stdLongWeekDay:
			s := Weekday((int(j.Weekday()) + 1) % 7).String()
			b = append(b, s...)
		case stdDay:
			b = appendInt(b, day, 0)
		case stdUnderDay:
			if day < 10 {
				b = append(b, ' ')
			}
			b = appendInt(b, day, 0)
		case stdZeroDay:
			b = appendInt(b, day, 2)
		case stdHour:
			b = appendInt(b, hour, 2)
		case stdHour12:
			// Noon is 12PM, midnight is 12AM.
			hr := hour % 12
			if hr == 0 {
				hr = 12
			}
			b = appendInt(b, hr, 0)
		case stdZeroHour12:
			// Noon is 12PM, midnight is 12AM.
			hr := hour % 12
			if hr == 0 {
				hr = 12
			}
			b = appendInt(b, hr, 2)
		case stdMinute:
			b = appendInt(b, min, 0)
		case stdZeroMinute:
			b = appendInt(b, min, 2)
		case stdSecond:
			b = appendInt(b, sec, 0)
		case stdZeroSecond:
			b = appendInt(b, sec, 2)
		case stdPM, stdpm:
			if hour >= 12 {
				b = append(b, "بعدازظهر"...)
			} else {
				b = append(b, "قبل‌ازظهر"...)
			}
		case stdFracSecond0, stdFracSecond9:
			b = formatNano(b, uint(j.Nanosecond()), std>>stdArgShift, std&stdMask == stdFracSecond9)
		}
	}
	return b, nil
}

// nextStdChunk finds the first occurrence of a std string in
// layout and returns the text before, the std string, and the text after.
func nextStdChunk(layout string) (prefix string, std int, suffix string) {
	for i := 0; i < len(layout); i++ {
		switch c := int(layout[i]); c {
		case 'J': // January, Jan
			if len(layout) >= i+3 && layout[i:i+3] == "Jan" {
				if len(layout) >= i+7 && layout[i:i+7] == "January" {
					return layout[0:i], stdLongMonth, layout[i+7:]
				}
				if !startsWithLowerCase(layout[i+3:]) {
					return layout[0:i], stdMonth, layout[i+3:]
				}
			}

		case 'M': // Monday, Mon
			if layout[i:i+3] == "Mon" {
				if len(layout) >= i+6 && layout[i:i+6] == "Monday" {
					return layout[0:i], stdLongWeekDay, layout[i+6:]
				}
				if !startsWithLowerCase(layout[i+3:]) {
					return layout[0:i], stdWeekDay, layout[i+3:]
				}
			}

		case '0': // 01, 02, 03, 04, 05, 06
			if len(layout) >= i+2 && '1' <= layout[i+1] && layout[i+1] <= '6' {
				return layout[0:i], std0x[layout[i+1]-'1'], layout[i+2:]
			}

		case '1': // 15, 1
			if len(layout) >= i+2 && layout[i+1] == '5' {
				return layout[0:i], stdHour, layout[i+2:]
			}
			return layout[0:i], stdNumMonth, layout[i+1:]

		case '2': // 2006, 2
			if len(layout) >= i+4 && layout[i:i+4] == "2006" {
				return layout[0:i], stdLongYear, layout[i+4:]
			}
			return layout[0:i], stdDay, layout[i+1:]

		case '_': // _2, _2006
			if len(layout) >= i+2 && layout[i+1] == '2' {
				//_2006 is really a literal _, followed by stdLongYear
				if len(layout) >= i+5 && layout[i+1:i+5] == "2006" {
					return layout[0 : i+1], stdLongYear, layout[i+5:]
				}
				return layout[0:i], stdUnderDay, layout[i+2:]
			}

		case '3':
			return layout[0:i], stdHour12, layout[i+1:]

		case '4':
			return layout[0:i], stdMinute, layout[i+1:]

		case '5':
			return layout[0:i], stdSecond, layout[i+1:]

		case 'P': // PM
			if len(layout) >= i+2 && layout[i+1] == 'M' {
				return layout[0:i], stdPM, layout[i+2:]
			}

		case 'p': // pm
			if len(layout) >= i+2 && layout[i+1] == 'm' {
				return layout[0:i], stdpm, layout[i+2:]
			}

		case '.': // .000 or .999 - repeated digits for fractional seconds.
			if i+1 < len(layout) && (layout[i+1] == '0' || layout[i+1] == '9') {
				ch := layout[i+1]
				j := i + 1
				for j < len(layout) && layout[j] == ch {
					j++
				}
				// String of digits must end here - only fractional second is all digits.
				if !isDigit(layout, j) {
					std := stdFracSecond0
					if layout[i+1] == '9' {
						std = stdFracSecond9
					}
					std |= (j - (i + 1)) << stdArgShift
					return layout[0:i], std, layout[j:]
				}
			}
		}
	}
	return layout, 0, ""
}

// startsWithLowerCase reports whether the string has a lower-case letter at the beginning.
// Its purpose is to prevent matching strings like "Month" when looking for "Mon".
func startsWithLowerCase(str string) bool {
	if len(str) == 0 {
		return false
	}
	c := str[0]
	return 'a' <= c && c <= 'z'
}

// isDigit reports whether s[i] is in range and is a decimal digit.
func isDigit(s string, i int) bool {
	if len(s) <= i {
		return false
	}
	c := s[i]
	return '0' <= c && c <= '9'
}

// appendInt appends the decimal form of x to b and returns the result.
// If the decimal form (excluding sign) is shorter than width, the result is padded with leading 0's.
// Duplicates functionality in strconv, but avoids dependency.
func appendInt(b []byte, x int, width int) []byte {
	u := uint(x)
	if x < 0 {
		b = append(b, '-')
		u = uint(-x)
	}

	// Assemble decimal in reverse order.
	var buf [20]rune
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = rune('۰' + u - q*10)
		u = q
	}
	i--
	buf[i] = rune('۰' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		b = append(b, []byte("۰")...)
	}

	return append(b, []byte(string(buf[i:]))...)
}

// formatNano appends a fractional second, as nanoseconds, to b
// and returns the result.
func formatNano(b []byte, nanosec uint, n int, trim bool) []byte {
	u := nanosec
	var buf [9]rune
	for start := len(buf); start > 0; {
		start--
		buf[start] = rune(u%10 + '۰')
		u /= 10
	}

	if n > 9 {
		n = 9
	}
	if trim {
		for n > 0 && buf[n-1] == '۰' {
			n--
		}
		if n == 0 {
			return b
		}
	}
	b = append(b, '.')
	return append(b, []byte(string(buf[:n]))...)
}
