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
			var err error
			cap.Text[i], err = strconv.Unquote(`"` + line + `"`)
			if err != nil {
				cap.Text[i] = strings.Replace(line, "\\n", "\n", -1)
			}
		}
	}
	return subtitle
}
