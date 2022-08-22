package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/scraper/custom"
	"github.com/kkyr/go-recipe/internal/scraper/test"
)

func TestNewMinimalistBakerScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.MinimalistBakerHost)

	scraper, err := custom.NewMinimalistBakerScraper(doc)
	if err != nil {
		t.Fatalf("unexpected err while initializing scraper: %v", err)
	}

	scraperTest := test.Scraper{
		Author:       "Minimalist Baker",
		Categories:   []string{"Dip"},
		CookTime:     20 * time.Minute,
		Cuisine:      []string{"Gluten-Free", "Mexican-Inspired", "Vegan"},
		Description:  "These flavorful, creamy refried lentils are infused with smoky spices and chilis and made in 1 pot in less than 30 minutes! Say hello to the perfect side for Mexican night.",
		ImageURL:     "https://minimalistbaker.com/wp-content/uploads/2019/03/Smoky-Refried-Lentils-SQUARE.jpg",
		Ingredients:  []string{"1 ½ cups cooked lentils", "2 Tbsp avocado or coconut oil ((if avoiding oil, sub water))", "1/2 cup diced shallot or white onion", "3 cloves garlic ((minced))", "1/4 tsp sea salt, plus more to taste", "2 tsp ground smoked paprika", "2 tsp ground cumin", "1 tsp dried Mexican oregano ((optional))", "1-2 tsp coconut sugar ((or maple syrup or organic brown sugar))", "1-2 small chipotle peppers in adobo sauce ((from a can // roughly chopped — adjust according to preferred spice level))", "1/3 cup vegetable broth, plus more as needed ((or store-bought — we like Imagine brand))", "1-2 Tbsp lime juice", "Chips (corn or plantain)", "Cilantro"},
		Instructions: []string{"If cooking your lentils from scratch (not reflected in cook time), do so now by bringing twice the amount of water (or vegetable broth) as lentils to a boil. Then add lentils and bring back to a boil. Reduce heat to a simmer and cook uncovered until tender — depending on type of lentil and if they were soaked, anywhere from 7-20 minutes. Then drain and set aside. If using canned lentils, simply drain, rinse gently, and proceed with recipe.", "Heat a large saucepan over medium heat. Once hot, add oil and onion. Sauté until tender, about 4 minutes, then add garlic and continue sautéing until both are slightly golden brown and tender. Season with a pinch of salt, stir, and proceed to next step.", "Add cooked (drained) lentils, sea salt, smoked paprika, cumin, oregano (optional), coconut sugar, chipotle peppers (starting with 1 small pepper and adding more to taste), and vegetable broth (starting with recommended amount and adding more to achieve desired consistency). Hold off on adding the lime juice until the end of cooking.", "Cover and simmer on low for 10 minutes, stirring occasionally. Add more vegetable broth as needed if it gets too dry.", "Mash lentils with a potato masher, the back of your spoon, or (our preferred method) scoop out two-thirds of the lentils and, in a small blender or food processor, blend into a loose purée for a creamier texture. Then add back to saucepan and stir to incorporate.", "Add lime juice and stir. Then taste and adjust flavor as needed, adding more lime for acidity, paprika for smokiness, salt to taste, coconut sugar for sweetness, vegetable broth to thin, or chipotle peppers for heat.", "Serve with plantain or corn chips, or add to tacos, burritos, and more! Store leftovers covered in the refrigerator up to 4-5 days or in the freezer up to 1 month. Serve warm or cold. Reheat in a microwave or on the stovetop until hot, adding more vegetable broth or water to rehydrate as needed."},
		Language:     "",
		Name:         "Smoky 1-Pot Refried Lentils",
		Nutrition: recipe.Nutrition{
			Calories:              178,
			CarbohydrateGrams:     21.6,
			CholesterolMilligrams: 0,
			FatGrams:              7.5,
			FiberGrams:            5.5,
			ProteinGrams:          7.9,
			SaturatedFatGrams:     0.9,
			ServingSize:           "1 half-cup servings",
			SodiumMilligrams:      210,
			SugarGrams:            3.3,
			TransFatGrams:         0,
			UnsaturatedFatGrams:   0,
		},
		PrepTime:      10 * time.Minute,
		SuitableDiets: nil,
		TotalTime:     30 * time.Minute,
		Yields:        "4 (1/2-cup servings)",
	}

	scraperTest.Run(t, scraper)
}
