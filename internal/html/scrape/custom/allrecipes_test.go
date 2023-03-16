package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewAllRecipesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.AllRecipesHost)

	scraper, err := custom.NewAllRecipesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:        "Jackie M",
		Categories:    nil,
		CookTime:      30 * time.Minute,
		Cuisine:       nil,
		Description:   "This marinara sauce recipe is made by blending tomatoes, parsley, garlic, and oregano before simmering with onion and white wine for amazing flavor.",
		ImageURL:      "https://www.allrecipes.com/thmb/rAJNjIWA7FHaI4bveYdZFkCJ7oM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/11966-best-marinara-sauce-ddmfs-137-1x1-1-a9d124d5e86d463d87fa99de61bb9f02.jpg",
		Ingredients:   []string{"2 (14.5 ounce) cans stewed tomatoes", "1 (6 ounce) can tomato paste", "4 tablespoons chopped fresh parsley", "1 clove garlic, minced", "1 teaspoon dried oregano", "1 teaspoon salt", "0.25 teaspoon ground black pepper", "6 tablespoons olive oil", "0.33333334326744 cup finely diced onion", "0.5 cup white wine"},
		Instructions:  []string{"Place tomatoes, tomato paste, parsley, garlic, oregano, salt, and pepper in a food processor; blend until smooth.", "Heat oil in a large skillet over medium heat. Add onion and cook until slightly softened, about 2 minutes. Stir in blended tomato sauce and white wine.", "Simmer, stirring occasionally, until thickened, about 30 minutes."},
		Language:      "",
		Name:          "Best Marinara Sauce Yet",
		Nutrition:     recipe.Nutrition{Calories: 151, CarbohydrateGrams: 12, CholesterolMilligrams: 0, FatGrams: 11, FiberGrams: 2, ProteinGrams: 2, SaturatedFatGrams: 2, ServingSize: "", SodiumMilligrams: 685, SugarGrams: 7, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		PrepTime:      15 * time.Minute,
		SuitableDiets: nil,
		TotalTime:     45 * time.Minute,
		Yields:        "8",
	}

	scraperTest.Run(t, scraper)
}
