package subtitles

import (
	"fmt"
)

// FilterCaptions pass the captions through a filter function
func (subtitle *Subtitle) FilterCaptions(filter string) {
	switch filter {
	case "all":
		subtitle.filterCapitalization()
		subtitle.filterHTML()
		subtitle.filterOCR()
		subtitle.filterStyle()
	case "caps":
		subtitle.filterCapitalization()
	case "html":
		subtitle.filterHTML()
	case "ocr":
		subtitle.filterOCR()
	case "style":
		subtitle.filterStyle()
	case "none":
	default:
		fmt.Printf("Unrecognized filter name: %s\n", filter)
	}
	subtitle.cleanEmptyLines()
}

func (subtitle *Subtitle) cleanEmptyLines() {
	var ret []Caption
	for _, cap := range subtitle.Captions {
		var lines []string
		for _, v := range cap.Text {
			if len(v) > 0 {
				lines = append(lines, v)
			}
		}
		if len(lines) > 0 {
			cap.Text = lines
			ret = append(ret, cap)
		}
	}
	subtitle.Captions = ret
}
