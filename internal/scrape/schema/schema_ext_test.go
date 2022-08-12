package schema_test

import (
	"os"
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/scrape/schema"
	"github.com/kkyr/go-recipe/internal/scrape/test"

	"github.com/PuerkitoBio/goquery"
)

func TestGetRecipeScraper(t *testing.T) {
	file, err := os.Open("testdata/json-ld-schema.html")
	if err != nil {
		t.Fatalf("unexpected err while opening file: %v", err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatalf("unexpected err while reading file: %v", err)
	}

	scraperTest := test.Scraper{
		Author:       "Sara Buenfeld",
		Categories:   []string{"Dinner", "Lunch", "Supper"},
		CookTime:     5 * time.Minute,
		Cuisine:      nil,
		Description:  "Cook and serve this easy, healthy lunch in under 15 minutes. It gives you three of your 5-a-day, and the wholemeal noodles add fibre while colourful veg provides beta-carotene and vitamin C.",
		ImageURL:     "https://images.immediate.co.uk/production/volatile/sites/30/2020/08/hdp-noodle-salad-440-400-76177ff.jpg?resize=768,574",
		Ingredients:  []string{"1 tbsp sesame oil", "2 tsp tamari", "1 lemon , juiced", "1 red chilli , deseeded and finely chopped", "1 small onion , finely chopped", "2 wholemeal noodle nests (about 100g)", "160g sugar snap peas", "4 small clementines , peeled and chopped", "160g shredded carrots", "large handful of coriander , chopped", "50g roasted unsalted cashews"},
		Instructions: []string{"Mix all the dressing ingredients together in a large bowl, then stir in the onion. Meanwhile, cook the noodles in a pan of boiling water for 5 mins, adding the sugar snap peas halfway through the cooking time – the noodles and peas should be just tender. Drain, cool under cold running water and drain again. Snip or cut the noodles into smaller lengths to make them more manageable to eat.", "Tip the noodles and peas into the bowl with the dressing, along with the clementines, carrots, coriander and cashews. Toss to combine, then serve in bowls or pack into rigid airtight containers to take to work."},
		Language:     "",
		Name:         "Noodle salad with sesame dressing",
		Nutrition: recipe.Nutrition{
			Calories:              526,
			CarbohydrateGrams:     68,
			CholesterolMilligrams: 0,
			FatGrams:              19,
			FiberGrams:            11,
			ProteinGrams:          16,
			SaturatedFatGrams:     3,
			ServingSize:           "",
			SodiumMilligrams:      1.16,
			SugarGrams:            22,
			TransFatGrams:         0,
			UnsaturatedFatGrams:   0,
		},
		PrepTime:      7 * time.Minute,
		SuitableDiets: []recipe.Diet{recipe.GlutenFreeDiet, recipe.VeganDiet, recipe.VegetarianDiet},
		TotalTime:     12 * time.Minute,
		Yields:        "",
	}

	scraper, err := schema.GetRecipeScraper(doc)
	if err != nil {
		t.Fatalf("unexpected err while initializing scraper: %v", err)
	}

	scraperTest.Run(t, scraper)
}
