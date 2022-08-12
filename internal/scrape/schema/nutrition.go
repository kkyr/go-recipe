package schema

import (
	"regexp"
	"strconv"

	"github.com/kkyr/go-recipe"
)

var matchFirstNumber = regexp.MustCompile(`(\d*[.])?\d+`)

// ParseNutritionalInformation parses nutritional information (https://schema.org/NutritionInformation)
// that's encoded in a map into a Nutrition struct.
func ParseNutritionalInformation(data map[string]any) recipe.Nutrition {
	parseFloat := func(d any) float32 {
		if s, ok := d.(string); ok {
			d, _ := strconv.ParseFloat(matchFirstNumber.FindString(s), 32)
			return float32(d)
		}

		return 0
	}

	var servingSize string
	if s, ok := data["servingSize"].(string); ok {
		servingSize = s
	}

	// units are determined by schema format https://schema.org/NutritionInformation
	nutrition := recipe.Nutrition{
		Calories:              parseFloat(data["calories"]),
		CarbohydrateGrams:     parseFloat(data["carbohydrateContent"]),
		CholesterolMilligrams: parseFloat(data["cholesterolContent"]),
		FatGrams:              parseFloat(data["fatContent"]),
		FiberGrams:            parseFloat(data["fiberContent"]),
		ProteinGrams:          parseFloat(data["proteinContent"]),
		SaturatedFatGrams:     parseFloat(data["saturatedFatContent"]),
		ServingSize:           servingSize,
		SodiumMilligrams:      parseFloat(data["sodiumContent"]),
		SugarGrams:            parseFloat(data["sugarContent"]),
		TransFatGrams:         parseFloat(data["transFatContent"]),
		UnsaturatedFatGrams:   parseFloat(data["unsaturatedFatContent"]),
	}

	return nutrition
}
