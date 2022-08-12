package url

import (
	"regexp"
	"strings"
)

var urlRegExp = regexp.MustCompile(
	`^((?P<schema>.+?)://)?(?P<host>.*?)(:(?P<port>\d+?))?(?P<path>/.*?)?(?P<query>[?].*?)?$`)

// GetHost tries to determine the host from the passed url and returns it.
// The returned host will be stripped of the "www." subdomain, if it exists.
// e.g. GetHost("https://www.api.example.com") --> "api.example.com".
func GetHost(url string) string {
	matches := urlRegExp.FindStringSubmatch(url)
	return strings.TrimPrefix(matches[3], "www.")
}
