package subtitles

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Parse tries to parse a subtitle
func Parse(s string) (sub Subtitle, err error) {
	defer func() {
		if r := recover(); r != nil {
			sub = Subtitle{}
			err = fmt.Errorf("parse: unrecognized subtitle type")
		}
	}()
	if looksLikeSRT(s) {
		return NewFromSRT(s)
	} else if looksLikeSSA(s) {
		return NewFromSSA(s)
	} else if looksLikeCCDBCapture(s) {
		return NewFromCCDBCapture(s)
	} else if looksLikeDCSub(s) {
		return NewFromDCSub(s)
	}
	return Subtitle{}, fmt.Errorf("parse: unrecognized subtitle type")
}

// LooksLikeTextSubtitle returns true i byte stream seems to be of a recognized format
func LooksLikeTextSubtitle(filename string) bool {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := ConvertToUTF8(data)
	return looksLikeCCDBCapture(s) || looksLikeSSA(s) || looksLikeDCSub(s) || looksLikeSRT(s)
}
