package subtitles

import (
	"fmt"
	"io/ioutil"
	"log"

	"gonum.org/v1/gonum/stat"
)

// Parse tries to parse a subtitle
func Parse(s string) (sub Subtitle, err error) {
	var maxVariance float64
	key := -1
	var res [4]Subtitle
	res[0], _ = NewFromMicroDVD(s, -1)
	res[1], _ = NewFromSRT(s)
	res[2], _ = NewFromSSA(s)
	res[3], _ = NewFromTMPlayerTXT(s)

	// TODO: make the following convertion safe (will not panic
	// from any input), and then bring them back
	// res[4], _ = NewFromCCDBCapture(s)
	// res[5], _ = NewFromDCSub(s)

	// return the most varianced result
	for k, v := range res {
		vc := stat.Variance(toDataSet(v.AsPlainTXT()), nil)
		if vc > maxVariance {
			maxVariance = vc
			key = k
		}
	}
	if key >= 0 {
		return res[key], nil
	}
	return Subtitle{}, fmt.Errorf("parse: unrecognized subtitle type")
}

func toDataSet(s string) (res []float64) {
	for _, c := range s {
		res = append(res, float64(c))
	}
	return
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
