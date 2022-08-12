package recipe

import (
	recipe "go-recipe"
	"go-recipe/internal/scrape/custom"

	"github.com/PuerkitoBio/goquery"
)

type recipeScraperFunc func(*goquery.Document) (recipe.Scraper, error)

var hostToScraper = map[string]recipeScraperFunc{
	custom.ForksOverKnivesHost: custom.NewForksOverKnivesScraper,
	custom.MinimalistBakerHost: custom.NewMinimalistBakerScraper,
}
