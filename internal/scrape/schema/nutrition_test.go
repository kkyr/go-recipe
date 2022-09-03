package schema_test

import (
	"testing"

	"github.com/kkyr/go-recipe/internal/scrape/schema"

	"github.com/kkyr/assert"
)

func TestParseNutritionalInformation(t *testing.T) {
	assert := assert.New(t)

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

	assert.Field("Calories").Equal(float32(30), nutrition.Calories)
	assert.Field("CarbohydrateGrams").Equal(float32(8), nutrition.CarbohydrateGrams)
	assert.Field("CholesterolMilligrams").Equal(float32(3.5), nutrition.CholesterolMilligrams)
	assert.Field("FatGrams").Equal(float32(11.5), nutrition.FatGrams)
	assert.Field("FiberGrams").Equal(float32(105), nutrition.FiberGrams)
	assert.Field("ProteinGrams").Equal(float32(31), nutrition.ProteinGrams)
	assert.Field("SaturatedFatGrams").Equal(float32(201), nutrition.SaturatedFatGrams)
	assert.Field("ServingSize").Equal("4 loafs", nutrition.ServingSize)
	assert.Field("SodiumMilligrams").Equal(float32(78), nutrition.SodiumMilligrams)
	assert.Field("SugarGrams").Equal(float32(13.3), nutrition.SugarGrams)
	assert.Field("TransFatGrams").Equal(float32(0.55), nutrition.TransFatGrams)
	assert.Field("UnsaturatedFatGrams").Equal(float32(1.1), nutrition.UnsaturatedFatGrams)
}
