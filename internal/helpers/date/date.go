package date

import (
	"fmt"
	"time"
)

func ParseDate(date string) time.Time {
	layout := "2006-01-02-15-04-05"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t
}

func ParseHTMLDate(date string) time.Time {
	layout := "2006-01-02T15:04"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t
}

func GetTimestampFromDate(date time.Time) int64 {
	return date.Unix()
}

func GetDateFromTimestamp(date int64) time.Time {
	return time.Unix(date, 0).UTC()
}

func GetStringFromDate(date time.Time) string {
	return date.Format("2006-01-02-15-04-05")
}
