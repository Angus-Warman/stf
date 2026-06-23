package stf

import (
	"strings"
	"time"
)

var replacer = strings.NewReplacer(
	"yyyy", "2006",
	"fffffff", "0000000",
	"ffffff", "000000",
	"fffff", "00000",
	"ffff", "0000",
	"fff", "000",
	"yy", "06",
	"ff", "00",
	"f", "0",
	"zzz", "-07:00",
	"zz", "-0700",
	"z", "-07",
	"MMMM", "January",
	"MMM", "Jan",
	"MM", "01",
	"M", "1",
	"dddd", "Monday",
	"ddd", "Mon",
	"dd", "02",
	"d", "2",
	"HH", "15",
	"hh", "03",
	"h", "3",
	"mm", "04",
	"m", "4",
	"ss", "05",
	"s", "5",
	"tt", "PM",
	"K", "Z07:00",
)

// TimeFormats converts a time value to a string, using a layout similar to "yyyy/MM/dd HH:mm:ss"
func TimeFormat(t time.Time, layout string) string {
	goLayout := replacer.Replace(layout)

	return t.Format(goLayout)
}
