package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewLoveAndOtherSpicesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.LoveAndOtherSpicesHost)

	scraper, err := custom.NewLoveAndOtherSpicesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:        "Farwin",
		Categories:    []string{"Main Course"},
		CookTime:      20 * time.Minute,
		Cuisine:       []string{"Sri Lankan"},
		Description:   "Easy red lentil dal curry flavored with coconut milk, curry leaves and spices. A vegan, gluten free and high protein meal.",
		ImageURL:      "https://www.loveandotherspices.com/wp-content/uploads/2012/12/RedLentilDa-Thumbnail.jpg",
		Ingredients:   []string{"1 cup red lentils", "2.5 cups water (divided)", "1 small onion (chopped)", "1 small tomato (chopped)", "1 tsp ginger (minced)", "3 cloves garlic (sliced)", "1 green chili (halved)", "2 sprigs curry leaves", "1 tsp turmeric powder", "1 tsp red chili powder/cayenne pepper", "1 tsp cumin powder", "1/2 tsp coriander powder", "1 tsp salt", "1/2 cup full fat coconut milk", "1 tbsp oil", "1 tsp mustard seeds", "1/2 tsp cumin seeds", "1 dried chili (broken to pieces)", "1 sprig curry leaves", "1/2 tsp red chili flakes ((optional))"},
		Instructions:  []string{"Rinse the red lentils until the water runs clear.", "Into a pot add red lentils, 1.5 cups of water, chopped onion and tomatoes, green chili, curry leaves, minced ginger, sliced garlic, turmeric powder, red chili powder, cumin powder, coriander powder and salt. Bring to boil.", "Turn the heat to medium, cover the pot and cook the lentils until all the water is absorbed and lentils are cooked through. This would take about 15 minutes.", "Add the coconut milk and 1 cup of water, stir to combine. Cover and let it simmer for 5 minutes. Check seasoning and add more salt if needed.", "In a non-stick pan, heat the oil and add mustard seeds, cumin seeds, curry leaves, broken dried red chili. Cook until the mustard seeds start to splutter. Add red chili flakes if using.", "Pour the tempering to the cooked dal.", "Serve dal with rice, roti or your preferred carb."},
		Language:      "",
		Name:          "Sri Lankan Red lentil dal",
		Nutrition:     recipe.Nutrition{Calories: 0, CarbohydrateGrams: 0, CholesterolMilligrams: 0, FatGrams: 0, FiberGrams: 0, ProteinGrams: 0, SaturatedFatGrams: 0, ServingSize: "", SodiumMilligrams: 0, SugarGrams: 0, TransFatGrams: 0, UnsaturatedFatGrams: 0},
		PrepTime:      10 * time.Minute,
		SuitableDiets: nil,
		TotalTime:     30 * time.Minute,
		Yields:        "4",
	}

	scraperTest.Run(t, scraper)
}
