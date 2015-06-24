package engine1

import (
	"regexp"
)
func ParseError(err error) (code, text, index string) {
	re := regexp.MustCompile("^\\s*(.+)\\s*:\\s*(.+)\\s*\\[(.+)\\]\\s*$")

	if re.MatchString(err.Error()) {
		result := re.FindStringSubmatch(err.Error())
		code = result[1]
		text = result[2]
		index = result[3]
		return
	}
	return "", "", ""
}