package stf

import (
	"testing"
	"time"
)

func mustLoadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}

func TestTimeFormat(t *testing.T) {
	loc := mustLoadLocation("Australia/Sydney")
	tm := time.Date(2026, 6, 23, 14, 5, 7, 123000000, loc)

	tests := []struct {
		format   string
		expected string
	}{
		{"dd/MM/yyyy HH:mm:ss", "23/06/2026 14:05:07"},
		{"dd/MM/yyyy hh:mm:ss tt", "23/06/2026 02:05:07 PM"},
		{"yyyy-MM-dd", "2026-06-23"},
		{"HH:mm:ss.fff", "14:05:07.123"},
		{"dd MMM yyyy", "23 Jun 2026"},
		{"HH:mm:ss.ffffff", "14:05:07.123000"},
		{"HH:mm:ss.fffffff", "14:05:07.1230000"},
		{"yyyy-MM-ddTHH:mm:ss", "2026-06-23T14:05:07"},
		{"yy/MM/dd", "26/06/23"},
		{"h:mm tt", "2:05 PM"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			got := TimeFormat(tm, tt.format)
			if got != tt.expected {
				t.Errorf("TimeFormat(%v, %q) = %q, want %q", tm, tt.format, got, tt.expected)
			}
		})
	}
}

func TestTimeFormatTZ(t *testing.T) {
	loc := mustLoadLocation("America/New_York")
	tm := time.Date(2026, 6, 23, 8, 0, 0, 0, loc)

	tests := []struct {
		format   string
		expected string
	}{
		{"HH:mm:ss zzz", "08:00:00 -04:00"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			got := TimeFormat(tm, tt.format)
			if got != tt.expected {
				t.Errorf("TimeFormat(%v, %q) = %q, want %q", tm, tt.format, got, tt.expected)
			}
		})
	}
}

func TestTimeFormatUTC(t *testing.T) {
	tm := time.Date(2026, 6, 23, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		format   string
		expected string
	}{
		{"yyyy-MM-ddTHH:mm:ssK", "2026-06-23T00:00:00Z"},
		{"HH:mm:ss zzz", "00:00:00 +00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			got := TimeFormat(tm, tt.format)
			if got != tt.expected {
				t.Errorf("TimeFormat(%v, %q) = %q, want %q", tm, tt.format, got, tt.expected)
			}
		})
	}
}
