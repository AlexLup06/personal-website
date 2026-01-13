package helpers

import (
	"fmt"
	"time"
)

func ParseDate(dateString string) string {
	parsedTime, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	formattedDate := parsedTime.Format("2. January 2006")
	return formattedDate
}
