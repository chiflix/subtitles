package subtitles

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Eol is the end of line characters to use when writing .srt data
var eol = "\n"

func init() {
	if runtime.GOOS == "windows" {
		eol = "\r\n"
	}
}

func looksLikeSRT(s string) bool {
	caps, err := NewFromSRT(s)
	if err == nil && len(caps.Captions) >= 1 {
		return true
	}
	return false
}

// NewFromSRT parses a .srt text into Subtitle, assumes s is a clean utf8 string
func NewFromSRT(s string) (res Subtitle, err error) {
	re := regexp.MustCompile(`([0-9]+:*[0-9]+:[0-9]+[\.,]+[0-9]+)\s+-->\s+([0-9]+:*[0-9]+:[0-9]+[\.,]+[0-9]+)`)
	lines := strings.Split(s, "\n")
	outSeq := 1

	for i, line := range lines {
		matches := re.FindStringSubmatch(stripSpaces(line))
		if len(matches) < 3 {
			line = strings.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			_, err := strconv.Atoi(line)
			if err == nil {
				// if the next line is timecode
				if i+1 < len(lines) &&
					len(re.FindStringSubmatch(stripSpaces(lines[i+1]))) >= 3 {
					// then skip this seq number
					continue
				}
			}
			// not time codes, so it may be text
			ll := len(res.Captions) - 1
			if ll >= 0 {
				res.Captions[ll].Text = append(res.Captions[ll].Text, line)
			}
			continue
		}
		var o Caption
		o.Seq = outSeq

		o.Start, err = parseTime(matches[1])
		if err != nil {
			err = fmt.Errorf("srt: start error at line %d: %v", i, err)
			break
		}

		o.End, err = parseTime(matches[2])
		if err != nil {
			err = fmt.Errorf("srt: end error at line %d: %v", i, err)
			break
		}

		if removeLastEmptyCaption(&res, &o) {
			outSeq--
		} else {
			res.Captions = append(res.Captions, o)
			outSeq++
		}
	}

	removeLastEmptyCaption(&res, nil)
	return
}

func stripSpaces(line string) (r string) {
	return strings.Map(func(r rune) rune {
		if unicode.In(r, unicode.Number, unicode.Symbol,
			unicode.Punct,
			unicode.Dash, unicode.White_Space) {
			return r
		}
		return -1
	}, line)
}

func removeLastEmptyCaption(res *Subtitle, o *Caption) (removed bool) {
	ll := len(res.Captions)
	if ll > 0 && len(res.Captions[ll-1].Text) == 0 {
		// remove the last caption if it was empty
		if o != nil {
			o.Seq--
			res.Captions[ll-1] = *o
		} else {
			res.Captions = res.Captions[:ll-1]
		}
		return true
	}
	return false
}

// AsSRT renders the sub in .srt format
func (subtitle *Subtitle) AsSRT() (res string) {
	for _, sub := range subtitle.Captions {
		res += sub.AsSRT()
	}
	return
}

// AsSRT renders the caption as srt
func (cap Caption) AsSRT() string {
	res := fmt.Sprintf("%d", cap.Seq) + eol +
		TimeSRT(cap.Start) + " --> " + TimeSRT(cap.End) + eol
	for _, line := range cap.Text {
		res += line + eol
	}
	return res + eol
}

// TimeSRT renders a timestamp for use in .srt
func TimeSRT(t time.Time) string {
	res := t.Format("15:04:05.000")
	return strings.Replace(res, ".", ",", 1)
}
