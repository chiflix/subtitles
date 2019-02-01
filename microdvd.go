package subtitles

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func looksLikeMicroDVD(s string) bool {
	caps, err := NewFromMicroDVD(s, -1)
	if err == nil && len(caps.Captions) >= 1 {
		return true
	}
	return false
}

// NewFromMicroDVD parses a .sub text into Subtitle, assumes s is a clean utf8 string
// set frameRate to -1 to detect framerate or guess
func NewFromMicroDVD(s string, frameRate float64) (res Subtitle, err error) {
	re := regexp.MustCompile(`^\{([0-9]+)\}\{([0-9]+)\}(.*)$`)
	lines := strings.Split(s, "\n")
	outSeq := 1

	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) < 4 {
			err = fmt.Errorf("microdvd: parse error at line %d (idx out of range)", i)
			continue
		}
		if i == 0 && frameRate <= 0 {
			fps, err := strconv.ParseFloat(matches[3], 64)
			if err != nil {
				// there is not the framerate
				// use 23.976 by default
				frameRate = 23.976
				err = fmt.Errorf("microdvd: no frame rate assisned, use 23.976 by default ")
			} else {
				frameRate = fps
				continue
			}
		}
		startFrame, err := strconv.Atoi(matches[1])
		if err != nil {
			err = fmt.Errorf("microdvd: parse error at line %d (start frame is not int)", i)
			continue
		}
		endFrame, err := strconv.Atoi(matches[2])
		if err != nil {
			err = fmt.Errorf("microdvd: parse error at line %d (end frame is not int)", i)
			continue
		}
		var o Caption
		o.Start = secondsToTime(float64(startFrame) / frameRate)
		o.End = secondsToTime(float64(endFrame) / frameRate)
		o.Text = strings.Split(matches[3], "|")
		o.Seq = outSeq
		if len(o.Text) > 0 {
			res.Captions = append(res.Captions, o)
			outSeq++
		}
	}

	return res, nil
}
