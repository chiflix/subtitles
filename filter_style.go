package subtitles

import (
	"regexp"
	"strconv"
	"strings"
)

// filterStyle removes all html tags and ssa codes from captions
func (subtitle *Subtitle) filterStyle() *Subtitle {
	re := regexp.MustCompile("(?:<[^<>]+>|{[^{}]+})")
	for _, cap := range subtitle.Captions {
		for i, line := range cap.Text {
			if re.MatchString(line) {
				line = re.ReplaceAllString(line, "")
			}
			if isMeaningless(line) {
				cap.Text[i] = ""
				continue
			}
			line = unescapeString(line)
			line = strings.TrimSpace(line)
			cap.Text[i] = line
		}
	}
	return subtitle
}

func isMeaningless(line string) bool {
	arr := strings.Split(line, " ")
	if len(arr) > 5 {
		isnumberorsinglechar := 0
		for _, v := range arr {
			if len(v) == 1 {
				isnumberorsinglechar++
				continue
			}
			if _, err := strconv.Atoi(v); err == nil {
				isnumberorsinglechar++
			}
		}
		if isnumberorsinglechar == len(arr) {
			return true
		}
	}
	return false
}

func unescapeString(line string) string {
	r, err := strconv.Unquote(`"` + line + `"`)
	if err != nil {
		return strings.Replace(line, "\\n", "\n", -1)
	}
	return r
}
