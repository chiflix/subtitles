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
func NewFromMicroDVD(s string, fps float64) (res Subtitle, errs []error) {
	var err error
	re := regexp.MustCompile(`[\{\[]+([0-9]+)[\}\]]+[\{\[]+([0-9]+)[\}\]]+(.*)$`)
	lines := strings.Split(s, "\n")
	outSeq := 1

	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) < 4 {
			err = fmt.Errorf("microdvd: parse error at line %d (idx out of range)", i)
			errs = append(errs, err)
			continue
		}
		if i == 0 && fps <= 0 {
			fps, err = strconv.ParseFloat(matches[3], 64)
			if err != nil {
				// there is no framerate here
				// use 23.976 by default
				fps = 23.976
				err = fmt.Errorf("microdvd: no frame rate assigned, use 23.976 by default ")
				errs = append(errs, err)
			} else {
				continue
			}
		}
		startFrame, err := strconv.Atoi(matches[1])
		if err != nil {
			err = fmt.Errorf("microdvd: parse error at line %d (start frame is not int)", i)
			errs = append(errs, err)
			continue
		}
		endFrame, err := strconv.Atoi(matches[2])
		if err != nil {
			err = fmt.Errorf("microdvd: parse error at line %d (end frame is not int)", i)
			errs = append(errs, err)
			continue
		}
		var o Caption
		o.Start = secondsToTime(float64(startFrame) / fps)
		o.End = secondsToTime(float64(endFrame) / fps)
		o.Text = strings.Split(matches[3], "|")
		o.Seq = outSeq
		if len(o.Text) > 0 {
			res.Captions = append(res.Captions, o)
			outSeq++
		}
	}

	return res, nil
}
