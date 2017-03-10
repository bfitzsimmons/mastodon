package mastodon

import (
	"fmt"
	"strings"
)

// GeneratePlaceholders returns the field placeholders for use in SQL queries.
func GeneratePlaceholders(num int) string {
	placeholders := ""
	for i := 0; i < num; i++ {
		placeholders += fmt.Sprintf("$%d, ", i+1)
	}
	return strings.TrimSuffix(placeholders, ", ")
}

// GenerateNamedPlaceholders returns named placeholders (:field_name) for use in SQL queries.
func GenerateNamedPlaceholders(fields []string) string {
	placeholders := ""
	for _, field := range fields {
		placeholders += fmt.Sprintf(":%s, ", field)
	}
	return strings.TrimSuffix(placeholders, ", ")
}
