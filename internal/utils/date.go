package utils

import (
	"strings"
	"time"
)

func FormatDate(updatedAt string) (string, error) {
	date, err := time.Parse(time.RFC3339, updatedAt)

	if err != nil {
		return "", err
	}

	// Get the local timezone of Argentina
	local, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return "", err
	}

	// Format the date with the local timezone
	formattedDate := date.In(local).Format(time.RFC3339)

	dateTimeSplited := strings.Split(formattedDate, "T")
	hours := strings.Split(dateTimeSplited[1], ":0")
	updatedAtFormatted := dateTimeSplited[0] + " " + hours[0]

	return updatedAtFormatted, nil
}
