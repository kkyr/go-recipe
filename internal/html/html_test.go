package html_test

import (
	"testing"

	"github.com/kkyr/go-recipe/internal/html"

	"github.com/kkyr/assert"
)

func TestCleanString(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		in   string
		want string
	}{
		{
			in:   "unchanged",
			want: "unchanged",
		},
		{
			in:   "1 &gt; 2",
			want: "1 > 2",
		},
		{
			in:   "\t\ttabs\n\n vs   spaces\t  ",
			want: "tabs vs spaces",
		},
		{
			in:   "It    was   a bright   day in    April  ",
			want: "It was a bright day in April",
		},
		{
			in:   "https://www.example.com/",
			want: "https://www.example.com/",
		},
	} {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(tc.want, html.CleanString(tc.in))
		})
	}
}
