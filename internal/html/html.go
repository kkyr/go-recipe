package html

import (
	"html"
	"strings"
)

// CleanString cleans s by un-escaping special HTML characters and eliminating extraneous whitespace.
func CleanString(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	s = html.UnescapeString(s)
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")

	return s
}
