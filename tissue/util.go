package tissue

import (
	"regexp"
	"strings"
)

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "-", -1)
	nonAlphaNumDash := regexp.MustCompile("[^A-Za-z0-9-]")
	s = nonAlphaNumDash.ReplaceAllLiteralString(s, "-")
	multiDash := regexp.MustCompile("-+")
	s = multiDash.ReplaceAllLiteralString(s, "-")
	endDash := regexp.MustCompile("-$")
	s = endDash.ReplaceAllLiteralString(s, "")
	return s
}
