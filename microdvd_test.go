package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromMicroDVD(t *testing.T) {
	in := `{1}{1}23.976
{0}{43}What's my favorite way to relax?
{103}{167}And disappear into|the nature channel.`

	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 0, 0),
		makeTime(0, 0, 1, 793),
		[]string{"What's my favorite way to relax?"},
	}, {
		2,
		makeTime(0, 0, 4, 296),
		makeTime(0, 0, 6, 965),
		[]string{"And disappear into",
			"the nature channel."},
	}}}

	res, err := NewFromMicroDVD(in, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, res)
}
