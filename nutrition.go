package recipe

// Nutrition represents nutritional information about a recipe.
type Nutrition struct {
	// The number of calories.
	Calories float32
	// The number of grams of carbohydrates.
	CarbohydrateGrams float32
	// The number of milligrams of cholesterol.
	CholesterolMilligrams float32
	// The number of grams of fat.
	FatGrams float32
	// The number of grams of fiber.
	FiberGrams float32
	// The number of grams of protein.
	ProteinGrams float32
	// The number of grams of saturated fat.
	SaturatedFatGrams float32
	// The serving size, in terms of the number of volume or mass.
	ServingSize string
	// The number of milligrams of sodium.
	SodiumMilligrams float32
	// The number of grams of sugar.
	SugarGrams float32
	// The number of grams of trans fat.
	TransFatGrams float32
	// The number of grams of unsaturated fat.
	UnsaturatedFatGrams float32
}
