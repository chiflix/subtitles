package subtitles

import (
	"regexp"
)

// filterStyle removes all html tags and ssa codes from captions
func (subtitle *Subtitle) filterStyle() *Subtitle {
	re := regexp.MustCompile("(?:<[^<>]+>|{[^{}]+})")
	for _, cap := range subtitle.Captions {
		for i, line := range cap.Text {
			if re.MatchString(line) {
				cap.Text[i] = re.ReplaceAllString(line, "")
			}
		}
	}
	return subtitle
}
