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
}
