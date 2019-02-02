package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromTMPlayerTXT(t *testing.T) {
	in := "00:01:02:-First line of text|-Second line of text\n" +
		"00:01:04:Another line\n" +
		"02:44:Another line ok\n"

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 1, 2, 0),
		makeTime(0, 1, 3, 999),
		[]string{"-First line of text",
			"-Second line of text"},
	}, {
		2,
		makeTime(0, 1, 4, 0),
		makeTime(0, 1, 7, 0),
		[]string{"Another line"},
	}, {
		3,
		makeTime(0, 2, 44, 0),
		makeTime(0, 2, 47, 0),
		[]string{"Another line ok"},
	}}}

	res, err := NewFromTMPlayerTXT(in)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}
