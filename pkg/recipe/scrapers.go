package recipe

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kkyr/go-recipe"
	custom2 "github.com/kkyr/go-recipe/internal/html/scrape/custom"
)

type recipeScraperFunc func(*goquery.Document) (recipe.Scraper, error)

var hostToScraper = map[string]recipeScraperFunc{
	custom2.ForksOverKnivesHost: custom2.NewForksOverKnivesScraper,
	custom2.MinimalistBakerHost: custom2.NewMinimalistBakerScraper,
}
