package subtitles

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

/* Extensions: .txt
TMPlayer is the simplest possible subtitle format.
There are 2 variants: with short time codes and long time codes
(more popular). Line breaks are created with | character. You
can only specify when the text should appear and not for long.
Most programs either display every line for 3 seconds or
calculate how long the text should be displayed based on it's
length.

00:01:02:-First line of text|-Second line of text
02:03:44:Another line
*/

func looksLikeTMPlayerTXT(s string) bool {
	caps, err := NewFromTMPlayerTXT(s)
	if err == nil && len(caps.Captions) >= 1 {
		return true
	}
	return false
}

// NewFromTMPlayerTXT ...
func NewFromTMPlayerTXT(s string) (res Subtitle, errs []error) {
	var err error
	re := regexp.MustCompile(`([0-9]+:[0-9]+:[0-9]+):(.*)`)
	lines := strings.Split(s, "\n")
	outSeq := 1

	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) < 3 {
			continue
		}

		var o Caption
		o.Seq = outSeq
		o.Start, err = parseTime(matches[1])
		if err != nil {
			err = fmt.Errorf("srt: start error at line %d: %v", i, err)
			errs = append(errs, err)
			continue
		}
		ll := len(res.Captions)
		if ll > 0 {
			ll--
			if o.Start.After(res.Captions[ll].Start) &&
				res.Captions[ll].End.After(o.Start) {
				res.Captions[ll].End = o.Start.Add(-time.Millisecond)
			}
		}

		o.Text = strings.Split(matches[2], "|")

		o.End = o.Start.Add(3 * time.Duration(len(o.Text)) *
			time.Second)

		res.Captions = append(res.Captions, o)
		outSeq++
	}
	return res, nil
}
