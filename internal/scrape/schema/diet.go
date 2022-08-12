package schema

import (
	"strings"

	"github.com/kkyr/go-recipe"
)

// ParseDiet parses a restricted diet (https://schema.org/RestrictedDiet) string into
// a Diet type.
func ParseDiet(s string) recipe.Diet {
	clean := func(diet string) string {
		diet = strings.TrimSpace(diet)
		diet = strings.ToLower(diet)
		diet = strings.TrimPrefix(diet, "https://schema.org/")
		diet = strings.TrimSuffix(diet, "diet")
		diet = strings.Join(strings.Fields(diet), "")

		return diet
	}

	var strToDiet = map[string]recipe.Diet{
		"diabetic":   recipe.DiabeticDiet,
		"glutenfree": recipe.GlutenFreeDiet,
		"halal":      recipe.HalalDiet,
		"hindu":      recipe.HinduDiet,
		"kosher":     recipe.KosherDiet,
		"lowcalorie": recipe.LowCalorieDiet,
		"lowfat":     recipe.LowFatDiet,
		"lowlactose": recipe.LowLactoseDiet,
		"lowsalt":    recipe.LowSaltDiet,
		"vegan":      recipe.VeganDiet,
		"vegetarian": recipe.VegetarianDiet,
	}

	if diet, ok := strToDiet[clean(s)]; ok {
		return diet
	}

	return recipe.UnknownDiet
}
