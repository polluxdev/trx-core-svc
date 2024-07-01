package helper

import (
	"regexp"
	"strings"
	"time"
)

func ToSnakeCase(str string) string {
	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	snake = strings.ReplaceAll(snake, "__", "_")

	return strings.ToLower(snake)
}

func FormatDateTime(datetime *time.Time, format string) *string {
	if datetime == nil {
		return nil
	}

	result := datetime.Format(format)

	return &result
}
