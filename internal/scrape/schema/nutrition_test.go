package schema_test

import (
	"testing"

	"github.com/kkyr/go-recipe/internal/scrape/schema"
)

func TestParseNutritionalInformation(t *testing.T) {
	data := map[string]any{
		"calories":              "30 kcal",
		"carbohydrateContent":   "8grams",
		"cholesterolContent":    "3.5",
		"fatContent":            "11.5 grams fat",
		"fiberContent":          "105 g",
		"proteinContent":        "31",
		"saturatedFatContent":   "201",
		"servingSize":           "4 loafs",
		"sodiumContent":         "78 mg",
		"sugarContent":          "13.3g sugar",
		"transFatContent":       "0.55 g",
		"unsaturatedFatContent": "1.1 g",
	}

	nutrition := schema.ParseNutritionalInformation(data)

	if want := float32(30); nutrition.Calories != want {
		t.Errorf("Calories = %f, want %f", nutrition.Calories, want)
	}

	if want := float32(8); nutrition.CarbohydrateGrams != want {
		t.Errorf("CarbohydrateGrams = %f, want %f", nutrition.CarbohydrateGrams, want)
	}

	if want := float32(3.5); nutrition.CholesterolMilligrams != want {
		t.Errorf("CholesterolMilligrams = %f, want %f", nutrition.CholesterolMilligrams, want)
	}

	if want := float32(11.5); nutrition.FatGrams != want {
		t.Errorf("FatGrams = %f, want %f", nutrition.FatGrams, want)
	}

	if want := float32(105); nutrition.FiberGrams != want {
		t.Errorf("FiberGrams = %f, want %f", nutrition.FiberGrams, want)
	}

	if want := float32(31); nutrition.ProteinGrams != want {
		t.Errorf("ProteinGrams = %f, want %f", nutrition.ProteinGrams, want)
	}

	if want := float32(201); nutrition.SaturatedFatGrams != want {
		t.Errorf("SaturatedFatGrams = %f, want %f", nutrition.SaturatedFatGrams, want)
	}

	if want := data["servingSize"]; nutrition.ServingSize != want {
		t.Errorf("ServingSize = %s, want %s", nutrition.ServingSize, want)
	}

	if want := float32(78); nutrition.SodiumMilligrams != want {
		t.Errorf("SodiumMilligrams = %f, want %f", nutrition.SodiumMilligrams, want)
	}

	if want := float32(13.3); nutrition.SugarGrams != want {
		t.Errorf("SugarGrams = %f, want %f", nutrition.SugarGrams, want)
	}

	if want := float32(0.55); nutrition.TransFatGrams != want {
		t.Errorf("SugarGrams = %f, want %f", nutrition.TransFatGrams, want)
	}

	if want := float32(1.1); nutrition.UnsaturatedFatGrams != want {
		t.Errorf("SugarGrams = %f, want %f", nutrition.SugarGrams, want)
	}
}
