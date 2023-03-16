package recipe

import (
	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"

	"github.com/PuerkitoBio/goquery"
)

type recipeScraperFunc func(*goquery.Document) (recipe.Scraper, error)

var hostToScraper = map[string]recipeScraperFunc{
	custom.AllRecipesHost:         custom.NewAllRecipesScraper,
	custom.ForksOverKnivesHost:    custom.NewForksOverKnivesScraper,
	custom.LoveAndOtherSpicesHost: custom.NewLoveAndOtherSpicesScraper,
	custom.MinimalistBakerHost:    custom.NewMinimalistBakerScraper,
}
