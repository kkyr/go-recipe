package schema_test

import (
	"testing"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/kkyr/assert"
)

func TestParseDiet(t *testing.T) {
	assert := assert.New(t)

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
			assert.Equal(tc.want, schema.ParseDiet(tc.in))
		})
	}
}
