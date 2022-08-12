package html_test

import (
	"testing"

	"github.com/kkyr/go-recipe/internal/html"
)

func TestCleanString(t *testing.T) {
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
			got := html.CleanString(tc.in)
			if tc.want != got {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
