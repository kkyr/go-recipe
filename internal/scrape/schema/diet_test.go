package schema_test

import (
	"go-recipe/internal/scrape/schema"
	"testing"

	recipe "go-recipe"
)

func TestParseDiet(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want recipe.Diet
	}{
		{in: "Low Salt", want: recipe.LowSaltDiet},
		{in: "DiabeticDiet", want: recipe.DiabeticDiet},
		{in: "https://schema.org/GlutenFreeDiet", want: recipe.GlutenFreeDiet},
		{in: "halalDiet", want: recipe.HalalDiet},
		{in: "hindu", want: recipe.HinduDiet},
		{in: "not-a-diet", want: recipe.UnknownDiet},
	} {
		t.Run(tc.in, func(t *testing.T) {
			got := schema.ParseDiet(tc.in)
			if tc.want != got {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
