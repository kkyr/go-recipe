package url_test

import (
	"testing"

	"go-recipe/internal/url"
)

func TestGetHost(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want string
	}{
		{
			"http://example1.com",
			"example1.com",
		},
		{
			in:   "https://www.example2.com",
			want: "example2.com",
		},
		{
			in:   "http://example3.example.com",
			want: "example3.example.com",
		},
		{
			in:   "https://www.example4.com/example",
			want: "example4.com",
		},
		{
			in:   "https://www.example5.com:433",
			want: "example5.com",
		},
		{
			in:   "http://example6.com?a=b",
			want: "example6.com",
		},
		{
			in:   "https://example7.co.uk",
			want: "example7.co.uk",
		},
		{
			in:   "example8.org/hello",
			want: "example8.org",
		},
		{
			in:   "food.example9.com",
			want: "food.example9.com",
		},
		{
			in:   "www.example10.com/myrecipe/cake/123?ref=newsletter",
			want: "example10.com",
		},
		{
			in:   "https://www.food.recipe.example11.com",
			want: "food.recipe.example11.com",
		},
	} {
		t.Run(tc.in, func(t *testing.T) {
			if got := url.GetHost(tc.in); tc.want != got {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
