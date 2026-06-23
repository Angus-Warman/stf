package stf

import (
	"strings"
	"time"
)

var replacer = strings.NewReplacer(
	"yyyy", "2006",
	"yy", "06",
	"MM", "01",
	"dd", "02",
	"HH", "15",
	"hh", "03",
	"mm", "04",
	"ss", "05",
	"tt", "PM",
)

func TimeFormat(t time.Time, format string) string {
	goFormat := replacer.Replace(format)

	return t.Format(goFormat)
}
