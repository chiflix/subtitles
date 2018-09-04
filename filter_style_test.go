package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterStyle(t *testing.T) {
	in := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"<font color=#FFE87C>GO NINJA!</font>", "{\\1c&H6FCCF9&}NINJA GO!"},
	}}}
	expected := Subtitle{[]Caption{{
		1,
		makeTime(0, 0, 4, 630),
		makeTime(0, 0, 6, 18),
		[]string{"GO NINJA!", "NINJA GO!"},
	}}}
	assert.Equal(t, &expected, in.filterStyle())
}
