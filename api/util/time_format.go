package util

import (
	"time"
)

const DATE_TIME_LAYOUT = "2006-01-02 15:04"
const DATE_LAYOUT = "2006-01-02"

func FormatDateTime(d time.Time) string {
	return d.Format(DATE_TIME_LAYOUT)
}

func ParseDateTime(dateTimeStr string) time.Time {
	d, _ := time.Parse(DATE_TIME_LAYOUT, dateTimeStr)

	return d
}

func FormatDate(d time.Time) string {
	return d.Format(DATE_LAYOUT)
}

func ParseDate(dateStr string) time.Time {
	d, _ := time.Parse(DATE_LAYOUT, dateStr)

	return d
}
