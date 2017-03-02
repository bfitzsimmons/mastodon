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
